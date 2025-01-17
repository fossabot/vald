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

// Package service
package service

import (
	"fmt"

	"github.com/vdaas/vald/internal/net/grpc"
)

type MetaOption func(m *meta) error

var (
	defaultMetaOpts = []MetaOption{}
)

func WithMetaAddr(addr string) MetaOption {
	return func(m *meta) error {
		m.addr = addr
		return nil
	}
}

func WithMetaHostPort(host string, port int) MetaOption {
	return func(m *meta) error {
		m.addr = fmt.Sprintf("%s:%d", host, port)
		return nil
	}
}

func WithMetaClient(client grpc.Client) MetaOption {
	return func(m *meta) error {
		if client != nil {
			m.client = client
		}
		return nil
	}
}
