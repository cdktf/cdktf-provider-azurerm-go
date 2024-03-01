// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package virtualnetworkgatewayconnection


type VirtualNetworkGatewayConnectionCustomBgpAddresses struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.94.0/docs/resources/virtual_network_gateway_connection#primary VirtualNetworkGatewayConnection#primary}.
	Primary *string `field:"required" json:"primary" yaml:"primary"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.94.0/docs/resources/virtual_network_gateway_connection#secondary VirtualNetworkGatewayConnection#secondary}.
	Secondary *string `field:"optional" json:"secondary" yaml:"secondary"`
}

