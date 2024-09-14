// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package virtualnetworkgateway


type VirtualNetworkGatewayVpnClientConfigurationVirtualNetworkGatewayClientConnection struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.2.0/docs/resources/virtual_network_gateway#address_prefixes VirtualNetworkGateway#address_prefixes}.
	AddressPrefixes *[]*string `field:"required" json:"addressPrefixes" yaml:"addressPrefixes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.2.0/docs/resources/virtual_network_gateway#name VirtualNetworkGateway#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.2.0/docs/resources/virtual_network_gateway#policy_group_names VirtualNetworkGateway#policy_group_names}.
	PolicyGroupNames *[]*string `field:"required" json:"policyGroupNames" yaml:"policyGroupNames"`
}

