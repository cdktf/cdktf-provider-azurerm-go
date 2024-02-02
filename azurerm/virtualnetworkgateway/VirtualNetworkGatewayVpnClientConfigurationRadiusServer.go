// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package virtualnetworkgateway


type VirtualNetworkGatewayVpnClientConfigurationRadiusServer struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.90.0/docs/resources/virtual_network_gateway#address VirtualNetworkGateway#address}.
	Address *string `field:"required" json:"address" yaml:"address"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.90.0/docs/resources/virtual_network_gateway#score VirtualNetworkGateway#score}.
	Score *float64 `field:"required" json:"score" yaml:"score"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.90.0/docs/resources/virtual_network_gateway#secret VirtualNetworkGateway#secret}.
	Secret *string `field:"required" json:"secret" yaml:"secret"`
}

