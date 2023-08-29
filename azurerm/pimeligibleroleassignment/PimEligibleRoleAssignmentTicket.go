// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pimeligibleroleassignment


type PimEligibleRoleAssignmentTicket struct {
	// The ticket number.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/pim_eligible_role_assignment#number PimEligibleRoleAssignment#number}
	Number *string `field:"optional" json:"number" yaml:"number"`
	// The ticket system.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/pim_eligible_role_assignment#system PimEligibleRoleAssignment#system}
	SystemAttribute *string `field:"optional" json:"systemAttribute" yaml:"systemAttribute"`
}

