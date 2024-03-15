// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

//go:build ipam_provider_volcengine

package cmd

import (
	"github.com/cilium/cilium/pkg/ipam/allocator/volcengine"
	ipamOption "github.com/cilium/cilium/pkg/ipam/option"
)

func init() {
	allocatorProviders[ipamOption.IPAMVolcengine] = new(volcengine.AllocatorVolcengine)
}
