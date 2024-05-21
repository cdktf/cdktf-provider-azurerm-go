// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package localnetworkgateway


type LocalNetworkGatewayTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.2/docs/resources/local_network_gateway#create LocalNetworkGateway#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.2/docs/resources/local_network_gateway#delete LocalNetworkGateway#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.2/docs/resources/local_network_gateway#read LocalNetworkGateway#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.2/docs/resources/local_network_gateway#update LocalNetworkGateway#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

