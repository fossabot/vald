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

package service

import (
	"context"
	"reflect"
	"unsafe"

	"github.com/vdaas/vald/internal/compress"
	"github.com/vdaas/vald/internal/errgroup"
	"github.com/vdaas/vald/internal/errors"
)

type Compressor interface {
	Compress(ctx context.Context, vector []float64) (string, error)
	Decompress(ctx context.Context, str string) ([]float64, error)
	MultiCompress(ctx context.Context, vectors [][]float64) ([]string, error)
	MultiDecompress(ctx context.Context, strs []string) ([][]float64, error)
}

type compressor struct {
	compressor compress.Compressor
	limitation int
}

func NewCompressor(opts ...CompressorOption) (Compressor, error) {
	c := new(compressor)
	for _, opt := range append(defaultCompressorOpts, opts...) {
		if err := opt(c); err != nil {
			return nil, errors.ErrOptionFailed(err, reflect.ValueOf(opt))
		}
	}

	return c, nil
}

func (c *compressor) Compress(ctx context.Context, vector []float64) (string, error) {
	res, err := c.compressor.CompressVector(vector)
	if err != nil {
		return "", err
	}

	return *(*string)(unsafe.Pointer(&res)), nil
}

func (c *compressor) Decompress(ctx context.Context, str string) ([]float64, error) {
	return c.compressor.DecompressVector(*(*[]byte)(unsafe.Pointer(&str)))
}

func (c *compressor) MultiCompress(ctx context.Context, vectors [][]float64) ([]string, error) {
	eg, ctx := errgroup.New(ctx)

	eg.Limitation(c.limitation)

	res := make([]string, len(vectors))
	for i, vector := range vectors {
		eg.Go(func() error {
			r, err := c.compressor.CompressVector(vector)
			if err != nil {
				return err
			}

			res[i] = *(*string)(unsafe.Pointer(&r))

			return nil
		})
	}

	err := eg.Wait()
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *compressor) MultiDecompress(ctx context.Context, strs []string) ([][]float64, error) {
	eg, ctx := errgroup.New(ctx)

	eg.Limitation(c.limitation)

	res := make([][]float64, len(strs))
	for i, str := range strs {
		eg.Go(func() error {
			r, err := c.compressor.DecompressVector(*(*[]byte)(unsafe.Pointer(&str)))
			if err != nil {
				return err
			}

			res[i] = r

			return nil
		})
	}

	err := eg.Wait()
	if err != nil {
		return nil, err
	}

	return res, nil
}
