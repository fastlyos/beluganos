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

	log "github.com/sirupsen/logrus"
)

//
// FIBCFlowMod process any FlowMod.
//
func (s *Server) FIBCFlowMod(hdr *fibcnet.Header, mod *fibcapi.FlowMod) {
	// s.log.Debugf("FlowMod: %v %v", hdr, mod)
	fibcapi.DispatchFlowMod(hdr, mod, s)
}

//
// FIBCMPLSFlowMod process FlowMod (MPLS)
//
func (s *Server) FIBCMPLSFlowMod(hdr *fibcnet.Header, mod *fibcapi.FlowMod, flow *fibcapi.MPLSFlow) {
	s.log.Debugf("FlowMod(MPLS): %v", hdr)
	fibcapi.LogFlowMod(s.log, log.DebugLevel, mod)
}
