// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package roleassignment


type RoleAssignmentTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.24.0/docs/resources/role_assignment#create RoleAssignment#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.24.0/docs/resources/role_assignment#delete RoleAssignment#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.24.0/docs/resources/role_assignment#read RoleAssignment#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}

