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

package mysql

import (
	dbr "github.com/gocraft/dbr/v2"
)

const (
	comma = ","
)

type MetaVector interface {
	GetUUID() string
	GetVector() string
	GetMeta() string
	GetIPs() []string
}

type metaVector struct {
	meta   meta
	podIPs []podIP
}

type meta struct {
	UUID   string         `db:"uuid"`
	Vector string         `db:"vector"`
	Meta   dbr.NullString `db:"meta"`
}

type podIP struct {
	UUID string `db:"uuid"`
	IP   string `db:"ip"`
}

func (m *metaVector) GetUUID() string   { return m.meta.UUID }
func (m *metaVector) GetVector() string { return m.meta.Vector }
func (m *metaVector) GetMeta() string   { return m.meta.Meta.String }
func (m *metaVector) GetIPs() []string {
	ips := make([]string, 0, len(m.podIPs))

	for _, ip := range m.podIPs {
		ips = append(ips, ip.IP)
	}

	return ips
}
