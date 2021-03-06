// -*- coding: utf-8 -*-

// Copyright (C) 2018 Nippon Telegraph and Telephone Corporation.
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

package gonslib

import (
	fibcapi "fabricflow/fibc/api"
	fibcnet "fabricflow/fibc/net"

	"github.com/beluganos/go-opennsl/opennsl"
	log "github.com/sirupsen/logrus"
)

//
// FIBCL3UnicastGroupMod process GroupMod(L3 Unicst)
//
func (s *Server) FIBCL3UnicastGroupMod(hdr *fibcnet.Header, mod *fibcapi.GroupMod, group *fibcapi.L3UnicastGroup) {
	s.log.Debugf("GroupMod(L3-UC): %v", hdr)
	fibcapi.LogGroupMod(s.log, log.DebugLevel, mod)

	neid := group.NeId
	vid := group.GetAdjustedVlanVid()
	_, portType := fibcapi.ParseDPPortId(group.PortId)

	var (
		trunk   opennsl.Trunk
		isTrunk bool
	)

	if portType == fibcapi.LinkType_BOND {
		if trunk, isTrunk = s.idmaps.Trunks.Get(group.PortId, vid); !isTrunk {
			s.log.Errorf("GroupMod(L3-UC): Trunk(port:%d, vid:%d) not found.", group.PortId, vid)
			return
		}
	}

	switch mod.Cmd {
	case fibcapi.GroupMod_ADD, fibcapi.GroupMod_MODIFY:
		oldL3egrID, ok := s.idmaps.L3Egress.Get(neid)
		if ok && (mod.Cmd != fibcapi.GroupMod_MODIFY) {
			s.log.Errorf("GroupMod(L3-UC): L3-UC(neid:%d) already exists. ", neid)
			return
		}

		ifaceID, ok := s.idmaps.L3Ifaces.Get(group.PortId, vid)
		if !ok {
			s.log.Errorf("GroupMod(L3-UC): L2-IF(port:%d, vid:%d) not found. ", group.PortId, vid)
			return
		}

		flags := opennsl.L3_NONE
		if mod.Cmd == fibcapi.GroupMod_MODIFY {
			flags = opennsl.L3_REPLACE | opennsl.L3_WITH_ID
		}

		phyPort := opennsl.Port(group.PhyPortId)
		pvid := s.vlanPorts.ConvVID(phyPort, opennsl.VlanDefaultMustGet(s.Unit()))

		tunnelInitiatorAdd(s.Unit(), group, ifaceID, pvid)
		tunnelTerminatorAdd(s.Unit(), group)

		l3egr := opennsl.NewL3Egress()
		l3egr.SetIfaceID(opennsl.L3IfaceID(ifaceID))
		l3egr.SetMAC(group.GetEthDstHwAddr())
		l3egr.SetVID(opennsl.Vlan(pvid))
		if isTrunk {
			flags |= opennsl.L3_TGID
			l3egr.SetTrunk(trunk)
		} else {
			l3egr.SetPort(phyPort)
		}

		l3egrID, err := l3egr.Create(s.Unit(), flags, oldL3egrID)
		if err != nil {
			s.log.Errorf("GroupMod(L3-UC): L3 Egress create error. %s", err)
			return
		}

		s.idmaps.L3Egress.Register(neid, l3egrID)

	case fibcapi.GroupMod_DELETE:
		l3egrID, ok := s.idmaps.L3Egress.Get(neid)
		if !ok {
			s.log.Errorf("GroupMod(L3-UC): L3-UC(%08x) not found. ", neid)
			return
		}

		s.idmaps.L3Egress.Unregister(neid)

		if err := l3egrID.Destroy(s.Unit()); err != nil {
			s.log.Errorf("GroupMod(L3-UC): L3 Egress delete error. %s", err)
		}

		tunnelTerminatorDelete(s.Unit(), group)

		vid := group.GetAdjustedVlanVid()
		ifaceID, ok := s.idmaps.L3Ifaces.Get(group.PortId, vid)
		if !ok {
			s.log.Errorf("GroupMod(L3-UC): L2-IF(port:%d, vid:%d) not found. ", group.PortId, vid)
			return
		}

		tunnelInitiatorDelete(s.Unit(), group, ifaceID)

	default:
		s.log.Errorf("GroupMod(L3-UC): Invalid Cmd. %d", mod.Cmd)
	}
}
