// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package virtualnetwork


type VirtualNetworkSubnetDelegation struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/virtual_network#name VirtualNetwork#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/virtual_network#service_delegation VirtualNetwork#service_delegation}.
	ServiceDelegation interface{} `field:"optional" json:"serviceDelegation" yaml:"serviceDelegation"`
}

