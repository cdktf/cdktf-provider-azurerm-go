// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package virtualnetworkgateway


type VirtualNetworkGatewayCustomRoute struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.28.0/docs/resources/virtual_network_gateway#address_prefixes VirtualNetworkGateway#address_prefixes}.
	AddressPrefixes *[]*string `field:"optional" json:"addressPrefixes" yaml:"addressPrefixes"`
}

