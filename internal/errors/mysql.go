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

// "github.com/pkg/errors"

var (
	// MySQL
	ErrMySQLConnectionPingFailed = New("error MySQL connection ping failed")

	NewErrMySQLNotFoundIdentity = func() error {
		return &ErrMySQLNotFoundIdentity{
			err: New("error mysql element not found"),
		}
	}

	ErrMySQLConnectionClosed = New("error MySQL connection closed")

	ErrRequiredElementNotFoundByUUID = func(uuid string) error {
		return Wrapf(NewErrMySQLNotFoundIdentity(), "error required element not found, uuid: %s", uuid)
	}

	NewErrMySQLInvalidArgumentIdentity = func() error {
		return &ErrMySQLInvalidArgumentIdentity{
			err: New("error mysql invalid argument"),
		}
	}

	ErrRequiredMemberNotFilled = func(member string) error {
		return Wrapf(NewErrMySQLInvalidArgumentIdentity(), "error required member not filled (member: %s)", member)
	}
)

type ErrMySQLNotFoundIdentity struct {
	err error
}

func (e *ErrMySQLNotFoundIdentity) Error() string {
	return e.err.Error()
}

func IsErrMySQLNotFound(err error) bool {
	switch err.(type) {
	case *ErrMySQLNotFoundIdentity:
		return true
	default:
		return false
	}
}

type ErrMySQLInvalidArgumentIdentity struct {
	err error
}

func (e *ErrMySQLInvalidArgumentIdentity) Error() string {
	return e.err.Error()
}

func IsErrMySQLInvalidArgument(err error) bool {
	switch err.(type) {
	case *ErrMySQLInvalidArgumentIdentity:
		return true
	default:
		return false
	}
}
