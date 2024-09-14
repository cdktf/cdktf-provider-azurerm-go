// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package stackhcilogicalnetwork


type StackHciLogicalNetworkSubnetRoute struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.2.0/docs/resources/stack_hci_logical_network#address_prefix StackHciLogicalNetwork#address_prefix}.
	AddressPrefix *string `field:"required" json:"addressPrefix" yaml:"addressPrefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.2.0/docs/resources/stack_hci_logical_network#name StackHciLogicalNetwork#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.2.0/docs/resources/stack_hci_logical_network#next_hop_ip_address StackHciLogicalNetwork#next_hop_ip_address}.
	NextHopIpAddress *string `field:"required" json:"nextHopIpAddress" yaml:"nextHopIpAddress"`
}

