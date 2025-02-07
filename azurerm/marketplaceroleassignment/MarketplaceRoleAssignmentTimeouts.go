// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package marketplaceroleassignment


type MarketplaceRoleAssignmentTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/marketplace_role_assignment#create MarketplaceRoleAssignment#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/marketplace_role_assignment#delete MarketplaceRoleAssignment#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/marketplace_role_assignment#read MarketplaceRoleAssignment#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}

