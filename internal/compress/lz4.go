//
// Copyright (C) 2019 Vdaas.org Vald team ( kpango, kou-m, rinx )
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

// Package compress provides compress functions
package compress

import (
	"bytes"
	"io/ioutil"
	"unsafe"

	"github.com/pierrec/lz4"
)

type lz4Compressor struct {
}

func NewLZ4() Compressor {
	return &lz4Compressor{}
}

func (l *lz4Compressor) CompressVector(vector []float64) ([]byte, error) {
	buf := new(bytes.Buffer)
	_, err := lz4.NewWriter(buf).Write(*(*[]byte)(unsafe.Pointer(&vector)))
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (l *lz4Compressor) DecompressVector(bs []byte) ([]float64, error) {
	r := lz4.NewReader(bytes.NewBuffer(bs))
	rawBytes, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	return *(*[]float64)(unsafe.Pointer(&rawBytes)), nil
}
