// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pimactiveroleassignment


type PimActiveRoleAssignmentTicket struct {
	// The ticket number.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/resources/pim_active_role_assignment#number PimActiveRoleAssignment#number}
	Number *string `field:"optional" json:"number" yaml:"number"`
	// The ticket system.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/resources/pim_active_role_assignment#system PimActiveRoleAssignment#system}
	SystemAttribute *string `field:"optional" json:"systemAttribute" yaml:"systemAttribute"`
}

