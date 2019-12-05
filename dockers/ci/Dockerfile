#
# Copyright (C) 2019 Vdaas.org Vald team ( kpango, kou-m, rinx )
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#    http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#

FROM vdaas/vald-base:latest AS builder

ARG BENCH_DATASET_MD5_DIR

WORKDIR ${GOPATH}/src/github.com/vdaas/vald

COPY Makefile .
COPY Makefile.d Makefile.d
COPY apis/proto apis/proto
COPY versions versions
COPY $BENCH_DATASET_MD5_DIR $BENCH_DATASET_MD5_DIR

RUN make ngt/install
RUN make bench/datasets

RUN apt-get update && apt-get install -y --no-install-recommends \
    libhdf5-dev \
    graphviz \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

RUN export LD_LIBRARY_PATH=$LD_LIBRARY_PATH:/usr/local/lib
