// -*- coding: utf-8 -*-

// Copyright (C) 2019 Nippon Telegraph and Telephone Corporation.
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

package mkpb

import "fmt"

func MakeSampleConfig(sampleName, lxcName string) (*Config, error) {
	switch sampleName {
	case "l3":
		return MakeSampleConfigL3(lxcName), nil
	case "l3-vlan":
		return MakeSampleConfigL3Vlan(lxcName), nil
	case "mpls-vpn":
		return MakeSampleConfigMplsVpn(lxcName), nil
	case "l2sw":
		return MakeSampleConfigL2SW(lxcName), nil
	case "iptun":
		return MakeSampleConfigIPTun(lxcName), nil
	default:
		return nil, fmt.Errorf("unknown sample. %s", sampleName)
	}
}

func MakeSampleConfigL3(lxcName string) *Config {
	config := NewConfig()
	config.Global.ReID = "1.1.1.1"
	config.Global.DpID = 1234
	config.Global.DpType = "as7712x4"
	config.Global.DpMode = "onsl"
	config.Global.DpAddr = "172.16.0.1"

	router := NewRouterConfig()
	router.Name = lxcName
	router.Daemons = []string{"zebra", "ospfd"}

	config.Router = []*RouterConfig{
		router,
	}

	return config
}

func MakeSampleConfigL3Vlan(lxcName string) *Config {
	config := NewConfig()
	config.Global.ReID = "1.1.1.1"
	config.Global.DpID = 1234
	config.Global.DpType = "as7712x4"
	config.Global.DpMode = "onsl"
	config.Global.DpAddr = "172.16.0.1"

	router := NewRouterConfig()
	router.Name = lxcName
	router.Eth = []uint32{1, 2, 3, 4, 5}
	router.Vlan = map[uint32][]uint16{
		3: []uint16{10},
		4: []uint16{10, 20},
		5: []uint16{20},
	}
	router.Daemons = []string{"zebra", "ospfd"}

	config.Router = []*RouterConfig{
		router,
	}

	return config
}

func MakeSampleConfigL2SW(lxcName string) *Config {
	config := NewConfig()
	config.Global.ReID = "1.1.1.1"
	config.Global.DpID = 1234
	config.Global.DpType = "as7712x4"
	config.Global.DpMode = "onsl"
	config.Global.DpAddr = "172.16.0.1"

	router := NewRouterConfig()
	router.Name = lxcName
	router.L2SW = NewL2SWConfig()
	router.L2SW.Access = map[uint32]uint16{
		1: 10,
		2: 20,
	}
	router.L2SW.Trunk = map[uint32][]uint16{
		3: []uint16{10, 20},
	}

	config.Router = []*RouterConfig{
		router,
	}

	return config
}

func MakeSampleConfigMplsVpn(lxcName string) *Config {
	config := NewConfig()
	config.Global.ReID = "1.1.1.1"
	config.Global.DpID = 1234
	config.Global.DpType = "as7712x4"
	config.Global.DpMode = "onsl"
	config.Global.DpAddr = "172.16.0.1"
	config.Global.Vpn = true

	mic := NewRouterConfig()
	mic.Name = lxcName
	mic.NodeID = 0
	mic.Eth = []uint32{1, 2, 3, 4, 5}
	mic.Daemons = []string{"zebra", "ospfd", "ldpd"}
	mic.RtRd = []string{"0:0", "0:0"}

	ric1 := NewRouterConfig()
	ric1.Name = "vpn-ric1"
	ric1.NodeID = 1
	ric1.Eth = []uint32{6, 7, 8, 9}
	ric1.Daemons = []string{"zebra"}
	ric1.RtRd = []string{"10:10", "10:101"}

	ric2 := NewRouterConfig()
	ric2.Name = "vpn-ric2"
	ric2.NodeID = 2
	ric2.Eth = []uint32{10, 11, 12, 13}
	ric2.Daemons = []string{"zebra"}
	ric2.RtRd = []string{"20:10", "20:101"}

	config.Router = []*RouterConfig{
		mic,
		ric1,
		ric2,
	}

	return config
}

func MakeSampleConfigIPTun(lxcName string) *Config {
	config := NewConfig()
	config.Global.ReID = "1.1.1.1"
	config.Global.DpID = 1234
	config.Global.DpType = "as7712x4"
	config.Global.DpMode = "onsl"
	config.Global.DpAddr = "172.16.0.1"

	router := NewRouterConfig()
	router.Name = lxcName
	router.Daemons = []string{"zebra", "ospfd", "ospfd6"}
	router.IPTun = MewIPTunConfig()
	router.IPTun.BgpRouteFamily = "ipv4-unicast"
	router.IPTun.LocalAddrRange4 = "10.0.0.0/24"
	router.IPTun.LocalAddrRange6 = "2001:2001::/64"
	router.IPTun.RemoteRoutes = []string{
		"10.1.1.0/24",
		"10.1.2.0/24",
		"2001:2010::/64",
		"2001:2020::/64",
	}

	config.Router = []*RouterConfig{
		router,
	}

	return config
}
