//
// Copyright (C) 2019 kpango (Yusuke Kato)
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

// Package server provides implementation of Go API for managing server flow
package server

import (
	"context"
	"crypto/tls"
	"net"
	"net/http"
	"reflect"
	"sync"
	"testing"
	"time"

	"github.com/vdaas/vald/internal/errgroup"
	"google.golang.org/grpc"
)

func TestNew(t *testing.T) {
	type args struct {
		opts []Option
	}
	tests := []struct {
		name    string
		args    args
		want    Server
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.opts...)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_server_IsRunning(t *testing.T) {
	type fields struct {
		mode mode
		name string
		mu   sync.RWMutex
		wg   sync.WaitGroup
		eg   errgroup.Group
		http struct {
			srv     *http.Server
			h       http.Handler
			starter func(net.Listener) error
		}
		grpc struct {
			srv  *grpc.Server
			opts []grpc.ServerOption
			reg  func(*grpc.Server)
		}
		l             net.Listener
		tcfg          *tls.Config
		pwt           time.Duration
		sddur         time.Duration
		rht           time.Duration
		rt            time.Duration
		wt            time.Duration
		it            time.Duration
		port          uint
		host          string
		enableRestart bool
		shuttingDown  bool
		running       bool
		preStartFunc  func() error
		preStopFunc   func() error
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &server{
				mode:          tt.fields.mode,
				name:          tt.fields.name,
				mu:            tt.fields.mu,
				wg:            tt.fields.wg,
				eg:            tt.fields.eg,
				http:          tt.fields.http,
				grpc:          tt.fields.grpc,
				l:             tt.fields.l,
				tcfg:          tt.fields.tcfg,
				pwt:           tt.fields.pwt,
				sddur:         tt.fields.sddur,
				rht:           tt.fields.rht,
				rt:            tt.fields.rt,
				wt:            tt.fields.wt,
				it:            tt.fields.it,
				port:          tt.fields.port,
				host:          tt.fields.host,
				enableRestart: tt.fields.enableRestart,
				shuttingDown:  tt.fields.shuttingDown,
				running:       tt.fields.running,
				preStartFunc:  tt.fields.preStartFunc,
				preStopFunc:   tt.fields.preStopFunc,
			}
			if got := s.IsRunning(); got != tt.want {
				t.Errorf("server.IsRunning() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_server_Name(t *testing.T) {
	type fields struct {
		mode mode
		name string
		mu   sync.RWMutex
		wg   sync.WaitGroup
		eg   errgroup.Group
		http struct {
			srv     *http.Server
			h       http.Handler
			starter func(net.Listener) error
		}
		grpc struct {
			srv  *grpc.Server
			opts []grpc.ServerOption
			reg  func(*grpc.Server)
		}
		l             net.Listener
		tcfg          *tls.Config
		pwt           time.Duration
		sddur         time.Duration
		rht           time.Duration
		rt            time.Duration
		wt            time.Duration
		it            time.Duration
		port          uint
		host          string
		enableRestart bool
		shuttingDown  bool
		running       bool
		preStartFunc  func() error
		preStopFunc   func() error
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &server{
				mode:          tt.fields.mode,
				name:          tt.fields.name,
				mu:            tt.fields.mu,
				wg:            tt.fields.wg,
				eg:            tt.fields.eg,
				http:          tt.fields.http,
				grpc:          tt.fields.grpc,
				l:             tt.fields.l,
				tcfg:          tt.fields.tcfg,
				pwt:           tt.fields.pwt,
				sddur:         tt.fields.sddur,
				rht:           tt.fields.rht,
				rt:            tt.fields.rt,
				wt:            tt.fields.wt,
				it:            tt.fields.it,
				port:          tt.fields.port,
				host:          tt.fields.host,
				enableRestart: tt.fields.enableRestart,
				shuttingDown:  tt.fields.shuttingDown,
				running:       tt.fields.running,
				preStartFunc:  tt.fields.preStartFunc,
				preStopFunc:   tt.fields.preStopFunc,
			}
			if got := s.Name(); got != tt.want {
				t.Errorf("server.Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_server_ListenAndServe(t *testing.T) {
	type fields struct {
		mode mode
		name string
		mu   sync.RWMutex
		wg   sync.WaitGroup
		eg   errgroup.Group
		http struct {
			srv     *http.Server
			h       http.Handler
			starter func(net.Listener) error
		}
		grpc struct {
			srv  *grpc.Server
			opts []grpc.ServerOption
			reg  func(*grpc.Server)
		}
		l             net.Listener
		tcfg          *tls.Config
		pwt           time.Duration
		sddur         time.Duration
		rht           time.Duration
		rt            time.Duration
		wt            time.Duration
		it            time.Duration
		port          uint
		host          string
		enableRestart bool
		shuttingDown  bool
		running       bool
		preStartFunc  func() error
		preStopFunc   func() error
	}
	type args struct {
		ech chan<- error
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &server{
				mode:          tt.fields.mode,
				name:          tt.fields.name,
				mu:            tt.fields.mu,
				wg:            tt.fields.wg,
				eg:            tt.fields.eg,
				http:          tt.fields.http,
				grpc:          tt.fields.grpc,
				l:             tt.fields.l,
				tcfg:          tt.fields.tcfg,
				pwt:           tt.fields.pwt,
				sddur:         tt.fields.sddur,
				rht:           tt.fields.rht,
				rt:            tt.fields.rt,
				wt:            tt.fields.wt,
				it:            tt.fields.it,
				port:          tt.fields.port,
				host:          tt.fields.host,
				enableRestart: tt.fields.enableRestart,
				shuttingDown:  tt.fields.shuttingDown,
				running:       tt.fields.running,
				preStartFunc:  tt.fields.preStartFunc,
				preStopFunc:   tt.fields.preStopFunc,
			}
			if err := s.ListenAndServe(tt.args.ech); (err != nil) != tt.wantErr {
				t.Errorf("server.ListenAndServe() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_server_Shutdown(t *testing.T) {
	type fields struct {
		mode mode
		name string
		mu   sync.RWMutex
		wg   sync.WaitGroup
		eg   errgroup.Group
		http struct {
			srv     *http.Server
			h       http.Handler
			starter func(net.Listener) error
		}
		grpc struct {
			srv  *grpc.Server
			opts []grpc.ServerOption
			reg  func(*grpc.Server)
		}
		l             net.Listener
		tcfg          *tls.Config
		pwt           time.Duration
		sddur         time.Duration
		rht           time.Duration
		rt            time.Duration
		wt            time.Duration
		it            time.Duration
		port          uint
		host          string
		enableRestart bool
		shuttingDown  bool
		running       bool
		preStartFunc  func() error
		preStopFunc   func() error
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &server{
				mode:          tt.fields.mode,
				name:          tt.fields.name,
				mu:            tt.fields.mu,
				wg:            tt.fields.wg,
				eg:            tt.fields.eg,
				http:          tt.fields.http,
				grpc:          tt.fields.grpc,
				l:             tt.fields.l,
				tcfg:          tt.fields.tcfg,
				pwt:           tt.fields.pwt,
				sddur:         tt.fields.sddur,
				rht:           tt.fields.rht,
				rt:            tt.fields.rt,
				wt:            tt.fields.wt,
				it:            tt.fields.it,
				port:          tt.fields.port,
				host:          tt.fields.host,
				enableRestart: tt.fields.enableRestart,
				shuttingDown:  tt.fields.shuttingDown,
				running:       tt.fields.running,
				preStartFunc:  tt.fields.preStartFunc,
				preStopFunc:   tt.fields.preStopFunc,
			}
			if err := s.Shutdown(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("server.Shutdown() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
