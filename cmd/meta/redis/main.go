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

// Package main provides program main
package main

import (
	"context"

	"github.com/vdaas/vald/internal/log"
	"github.com/vdaas/vald/internal/runner"
	"github.com/vdaas/vald/internal/safety"
	"github.com/vdaas/vald/pkg/meta/redis/config"
	"github.com/vdaas/vald/pkg/meta/redis/usecase"
)

const (
	// version represent the version
	version    = "v0.0.1"
	maxVersion = "v0.0.10"
	minVersion = "v0.0.0"
)

func main() {
	if err := safety.RecoverFunc(func() error {
		return runner.Do(
			context.Background(),
			runner.WithName("redis-meta"),
			runner.WithVersion(version, maxVersion, minVersion),
			runner.WithConfigLoader(func(path string) (interface{}, string, error) {
				cfg, err := config.NewConfig(path)
				if err != nil {
					return nil, "", err
				}
				return cfg, cfg.Version, err
			}),
			runner.WithDaemonInitializer(func(cfg interface{}) (runner.Runner, error) {
				return usecase.New(cfg.(*config.Data))
			}),
		)
	})(); err != nil {
		log.Fatal(err)
		return
	}
}
