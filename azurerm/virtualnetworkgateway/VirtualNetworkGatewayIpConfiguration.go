// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package virtualnetworkgateway


type VirtualNetworkGatewayIpConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/virtual_network_gateway#subnet_id VirtualNetworkGateway#subnet_id}.
	SubnetId *string `field:"required" json:"subnetId" yaml:"subnetId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/virtual_network_gateway#name VirtualNetworkGateway#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/virtual_network_gateway#private_ip_address_allocation VirtualNetworkGateway#private_ip_address_allocation}.
	PrivateIpAddressAllocation *string `field:"optional" json:"privateIpAddressAllocation" yaml:"privateIpAddressAllocation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/virtual_network_gateway#public_ip_address_id VirtualNetworkGateway#public_ip_address_id}.
	PublicIpAddressId *string `field:"optional" json:"publicIpAddressId" yaml:"publicIpAddressId"`
}

