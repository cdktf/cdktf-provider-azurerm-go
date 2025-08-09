// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pimeligibleroleassignment


type PimEligibleRoleAssignmentTicket struct {
	// User-supplied ticket number to be included with the request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.39.0/docs/resources/pim_eligible_role_assignment#number PimEligibleRoleAssignment#number}
	Number *string `field:"optional" json:"number" yaml:"number"`
	// User-supplied ticket system name to be included with the request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.39.0/docs/resources/pim_eligible_role_assignment#system PimEligibleRoleAssignment#system}
	SystemAttribute *string `field:"optional" json:"systemAttribute" yaml:"systemAttribute"`
}

