//
// Copyright (C) 2019 Vdaas.org Vald team ( kpango, kmrmt, rinx )
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package runner

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/vdaas/vald/internal/errgroup"
	"github.com/vdaas/vald/internal/errors"
	"github.com/vdaas/vald/internal/log"
	"github.com/vdaas/vald/internal/params"
	ver "github.com/vdaas/vald/internal/version"
)

type Runner interface {
	PreStart(ctx context.Context) error
	Start(ctx context.Context) <-chan error
	PreStop(ctx context.Context) error
	Stop(ctx context.Context) error
	PostStop(ctx context.Context) error
}

type runner struct {
	version          string
	maxVersion       string
	minVersion       string
	name             string
	loadConfig       func(string) (interface{}, string, error)
	initializeDaemon func(interface{}) (Runner, error)
}

func Do(ctx context.Context, opts ...Option) error {
	r := new(runner)

	for _, opt := range append(defaultOpts, opts...) {
		opt(r)
	}

	log.Init(log.DefaultGlg())

	p, isHelp, err := params.New(
		params.WithConfigFileDescription(fmt.Sprintf("%s config file path", r.name)),
	).Parse()

	if err != nil {
		return err
	}

	if isHelp {
		return nil
	}

	if p.ShowVersion() {
		log.Infof("vald %s server version -> %s", r.name, log.Bold(r.version))
		return nil
	}

	cfg, version, err := r.loadConfig(p.ConfigFilePath())
	if err != nil {
		return err
	}

	err = ver.Check(version, r.maxVersion, r.minVersion)
	if err != nil {
		return err
	}

	daemon, err := r.initializeDaemon(cfg)
	if err != nil {
		return err
	}

	log.Infof("service %s :%s starting...", r.name, version)

	return Run(ctx, daemon, r.name)
}

func Run(ctx context.Context, run Runner, name string) (err error) {
	sigCh := make(chan os.Signal, 1)
	defer close(sigCh)

	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	rctx, cancel := context.WithCancel(ctx)
	defer cancel()

	rctx = errgroup.Init(rctx)

	err = run.PreStart(rctx)
	if err != nil {
		return err
	}

	ech := run.Start(rctx)

	emap := make(map[string]int)
	errs := make([]error, 0, 10)

	for {
		select {
		case sig := <-sigCh:
			log.Warnf("%s signal received daemon will stopping soon...", sig)
			cancel()
		case err = <-ech:
			if err != nil {
				if _, ok := emap[err.Error()]; !ok {
					e := errors.ErrStartFunc(name, err)
					errs = append(errs, e)
					log.Error(err)
				}
				emap[err.Error()]++
			}
		case <-rctx.Done():
			err = run.PreStop(ctx)
			if err != nil {
				if _, ok := emap[err.Error()]; !ok {
					e := errors.ErrPreStopFunc(name, err)
					errs = append(errs, e)
					log.Error(err)
				}
				emap[err.Error()]++
			}
			err = run.Stop(ctx)
			if err != nil {
				if _, ok := emap[err.Error()]; !ok {
					e := errors.ErrStopFunc(name, err)
					errs = append(errs, e)
					log.Error(err)
				}
				emap[err.Error()]++
			}
			err = run.PostStop(ctx)
			if err != nil {
				if _, ok := emap[err.Error()]; !ok {
					e := errors.ErrPostStopFunc(name, err)
					errs = append(errs, e)
					log.Error(err)
				}
				emap[err.Error()]++
			}
			err = errgroup.Wait()
			if err != nil {
				if _, ok := emap[err.Error()]; !ok {
					e := errors.ErrRunnerWait(name, err)
					errs = append(errs, e)
					log.Error(err)
				}
				emap[err.Error()]++
			}
			err = nil
			for _, ierr := range errs {
				if ierr != nil {
					msg := ierr.Error()
					if msg != "" &&
						!strings.Contains(msg, http.ErrServerClosed.Error()) &&
						!strings.Contains(msg, context.Canceled.Error()) {
						err = errors.Wrapf(err, "error:\t%s\tcount:\t%d", msg, emap[msg])
					}
				}
			}
			if err != nil {
				err = errors.ErrDaemonStopFailed(err)
			}
			log.Warn("daemon stopped")
			return err
		}
	}
}
