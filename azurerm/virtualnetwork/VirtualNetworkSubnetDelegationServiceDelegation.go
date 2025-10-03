// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package virtualnetwork


type VirtualNetworkSubnetDelegationServiceDelegation struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.47.0/docs/resources/virtual_network#actions VirtualNetwork#actions}.
	Actions *[]*string `field:"optional" json:"actions" yaml:"actions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.47.0/docs/resources/virtual_network#name VirtualNetwork#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

