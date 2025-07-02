// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package userassignedidentity


type UserAssignedIdentityTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.35.0/docs/resources/user_assigned_identity#create UserAssignedIdentity#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.35.0/docs/resources/user_assigned_identity#delete UserAssignedIdentity#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.35.0/docs/resources/user_assigned_identity#read UserAssignedIdentity#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.35.0/docs/resources/user_assigned_identity#update UserAssignedIdentity#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

