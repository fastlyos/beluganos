// -*- coding: utf-8 -*-

// Copyright (C) 2017 Nippon Telegraph and Telephone Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ribsyn

import (
	"fabricflow/util/net"
	"net"
)

var (
	Tables *ribsTables
)

type ribsTables struct {
	Rics       *RicTable
	Nexthops   *NexthopTable
	NexthopMap *fflibnet.IPMap
}

func CreateTables(nexthops string) error {
	_, nw, err := net.ParseCIDR(nexthops)
	if err != nil {
		return err
	}

	ipgen := fflibnet.NewIPMapIPNetGenerator(nw)
	Tables = &ribsTables{
		Rics:       NewRicTable(),
		Nexthops:   NewNexthopTable(),
		NexthopMap: fflibnet.NewIPMap(ipgen),
	}
	return nil
}