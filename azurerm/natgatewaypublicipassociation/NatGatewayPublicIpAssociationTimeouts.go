// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package natgatewaypublicipassociation


type NatGatewayPublicIpAssociationTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.17.0/docs/resources/nat_gateway_public_ip_association#create NatGatewayPublicIpAssociation#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.17.0/docs/resources/nat_gateway_public_ip_association#delete NatGatewayPublicIpAssociation#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.17.0/docs/resources/nat_gateway_public_ip_association#read NatGatewayPublicIpAssociation#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}

