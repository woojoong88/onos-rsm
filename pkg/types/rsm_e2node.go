// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package types

import (
	"github.com/onosproject/onos-api/go/onos/topo"
)

type RsmE2Node struct {
	RsmE2NodeID string
	RsmNodeCapability []topo.RSMNodeSlicingCapabilityItem
}
