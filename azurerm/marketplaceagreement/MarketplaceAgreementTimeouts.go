// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package marketplaceagreement


type MarketplaceAgreementTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.87.0/docs/resources/marketplace_agreement#create MarketplaceAgreement#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.87.0/docs/resources/marketplace_agreement#delete MarketplaceAgreement#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.87.0/docs/resources/marketplace_agreement#read MarketplaceAgreement#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}

