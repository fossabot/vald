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

// Package errors provides error types and function
package errors

import "time"

// "github.com/pkg/errors"

var (
	// HTTP

	ErrInvalidAPIConfig = New("invalid api config")

	ErrInvalidRequest = New("invalid request")

	ErrHandler = func(err error) error {
		return Wrap(err, "handler returned error")
	}

	ErrHandlerTimeout = func(err error, t time.Time) error {
		return Wrapf(err, "handler timeout %v", t)
	}

	ErrRequestBodyCloseAndFlush = func(err error) error {
		return Wrap(err, "request body flush & close failed")
	}

	ErrRequestBodyClose = func(err error) error {
		return Wrap(err, "request body close failed")
	}

	ErrRequestBodyFlush = func(err error) error {
		return Wrap(err, "request body flush failed")
	}

	ErrTransportRetryable = New("transport is retryable")
)
