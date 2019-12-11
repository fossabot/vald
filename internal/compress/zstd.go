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
	"unsafe"

	// TODO
	// which is the better library of zstd algorithm?
	// "github.com/valyala/gozstd"
	"github.com/DataDog/zstd"
)

type zstdCompressor struct {
}

func NewZstd() Compressor {
	return &zstdCompressor{}
}

func (z *zstdCompressor) CompressVector(vector []float64) ([]byte, error) {
	return zstd.Compress(nil, *(*[]byte)(unsafe.Pointer(&vector)))
}

func (z *zstdCompressor) DecompressVector(bs []byte) ([]float64, error) {
	rawBytes, err := zstd.Decompress(nil, bs)
	if err != nil {
		return nil, err
	}

	return *(*[]float64)(unsafe.Pointer(&rawBytes)), nil
}
