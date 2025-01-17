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

import "runtime"

var (
	// Runtime

	ErrPanicRecovered = func(err error, rec interface{}) error {
		return Wrap(err, Errorf("panic recovered: %v", rec).Error())
	}

	ErrPanicString = func(err error, msg string) error {
		return Wrap(err, Errorf("panic recovered: %v", msg).Error())
	}

	ErrRuntimeError = func(err error, r runtime.Error) error {
		return Wrap(err, Errorf("system paniced caused by runtime error: %v", r).Error())
	}
)
