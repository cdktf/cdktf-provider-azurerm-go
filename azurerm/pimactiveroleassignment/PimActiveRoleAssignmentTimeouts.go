// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pimactiveroleassignment


type PimActiveRoleAssignmentTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/pim_active_role_assignment#create PimActiveRoleAssignment#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/pim_active_role_assignment#delete PimActiveRoleAssignment#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/pim_active_role_assignment#read PimActiveRoleAssignment#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}

