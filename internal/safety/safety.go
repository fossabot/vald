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

// Package safety provides safety functionality like revcover
package safety

import (
	"runtime"

	"github.com/vdaas/vald/internal/errors"
	"github.com/vdaas/vald/internal/log"
)

func RecoverFunc(fn func() error) func() error {
	return func() (err error) {
		defer func() {
			if r := recover(); r != nil {
				log.Warnf("recovered:\t%+v", r)
				switch x := r.(type) {
				case runtime.Error:
					err = errors.ErrRuntimeError(err, x)
					panic(err)
				case string:
					err = errors.ErrPanicString(err, x)
				case error:
					err = errors.Wrap(err, x.Error())
				default:
					err = errors.ErrPanicRecovered(err, x)
				}
				log.Error(err)
			}
		}()
		err = fn()
		return err
	}
}
