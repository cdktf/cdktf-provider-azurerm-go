// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package virtualnetworkgatewaynatrule


type VirtualNetworkGatewayNatRuleInternalMapping struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/virtual_network_gateway_nat_rule#address_space VirtualNetworkGatewayNatRule#address_space}.
	AddressSpace *string `field:"required" json:"addressSpace" yaml:"addressSpace"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/virtual_network_gateway_nat_rule#port_range VirtualNetworkGatewayNatRule#port_range}.
	PortRange *string `field:"optional" json:"portRange" yaml:"portRange"`
}

