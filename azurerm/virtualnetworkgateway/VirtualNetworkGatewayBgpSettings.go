// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package virtualnetworkgateway


type VirtualNetworkGatewayBgpSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.72.0/docs/resources/virtual_network_gateway#asn VirtualNetworkGateway#asn}.
	Asn *float64 `field:"optional" json:"asn" yaml:"asn"`
	// peering_addresses block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.72.0/docs/resources/virtual_network_gateway#peering_addresses VirtualNetworkGateway#peering_addresses}
	PeeringAddresses interface{} `field:"optional" json:"peeringAddresses" yaml:"peeringAddresses"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.72.0/docs/resources/virtual_network_gateway#peer_weight VirtualNetworkGateway#peer_weight}.
	PeerWeight *float64 `field:"optional" json:"peerWeight" yaml:"peerWeight"`
}

