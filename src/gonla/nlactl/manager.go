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

package nlactl

import (
	log "github.com/sirupsen/logrus"
	"gonla/nlamsg"
)

type NLAManager struct {
	NId   uint8
	Svcs  []NLAService
	Chans *NLAChannels
	done  <-chan struct{}
}

func NewNLAManager(nid uint8, done <-chan struct{}, svcs ...NLAService) *NLAManager {
	return &NLAManager{
		NId:   nid,
		Svcs:  svcs,
		Chans: NewNLAChannels(),
		done:  done,
	}
}

func (n *NLAManager) Add(svcs ...NLAService) {
	n.Svcs = append(n.Svcs, svcs...)
}

func (n *NLAManager) Serve() {
	defer n.Stop()

	for {
		select {
		case nlmsg := <-n.Chans.NlMsg:
			if err := nlamsg.Dispatch(nlmsg, n); err != nil {
				log.Warnf("NLAManager: Dispatch error. %s", err)
			}

		case nlmsg := <-n.Chans.Api:
			if err := nlamsg.DispatchUnion(nlmsg, n); err != nil {
				log.Warnf("NLAManager: DispatchUnion error. %s", err)
			}

		case <-n.done:
			log.Infof("NLAManager: EXIT")
			return
		}
	}
}

func (n *NLAManager) Start() error {
	for _, svc := range n.Svcs {
		if err := svc.Start(n.NId, n.Chans); err != nil {
			return err
		}
	}

	go n.Serve()

	log.Infof("NLAManager: START")
	return nil
}

func (n *NLAManager) Stop() {
	for _, svc := range n.Svcs {
		svc.Stop()
	}

	log.Infof("NLAManager: STOP")
}

func (n *NLAManager) NetlinkMessage(nlmsg *nlamsg.NetlinkMessage) {
	for _, svc := range n.Svcs {
		nlamsg.DispatchNetlinkMessage(nlmsg, svc)
	}
}

func (n *NLAManager) NetlinkAddr(nlmsg *nlamsg.NetlinkMessage, addr *nlamsg.Addr) {
	for _, svc := range n.Svcs {
		nlamsg.DispatchAddr(nlmsg, addr, svc)
	}
}

func (n *NLAManager) NetlinkLink(nlmsg *nlamsg.NetlinkMessage, link *nlamsg.Link) {
	for _, svc := range n.Svcs {
		nlamsg.DispatchLink(nlmsg, link, svc)
	}
}

func (n *NLAManager) NetlinkNeigh(nlmsg *nlamsg.NetlinkMessage, neigh *nlamsg.Neigh) {
	for _, svc := range n.Svcs {
		nlamsg.DispatchNeigh(nlmsg, neigh, svc)
	}
}

func (n *NLAManager) NetlinkRoute(nlmsg *nlamsg.NetlinkMessage, route *nlamsg.Route) {
	for _, svc := range n.Svcs {
		nlamsg.DispatchRoute(nlmsg, route, svc)
	}
}

func (n *NLAManager) NetlinkNode(nlmsg *nlamsg.NetlinkMessage, node *nlamsg.Node) {
	for _, svc := range n.Svcs {
		nlamsg.DispatchNode(nlmsg, node, svc)
	}
}

func (n *NLAManager) NetlinkVpn(nlmsg *nlamsg.NetlinkMessage, vpn *nlamsg.Vpn) {
	for _, svc := range n.Svcs {
		nlamsg.DispatchVpn(nlmsg, vpn, svc)
	}
}
