// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pimeligibleroleassignment


type PimEligibleRoleAssignmentScheduleExpiration struct {
	// The duration of the eligible role assignment in days.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.113.0/docs/resources/pim_eligible_role_assignment#duration_days PimEligibleRoleAssignment#duration_days}
	DurationDays *float64 `field:"optional" json:"durationDays" yaml:"durationDays"`
	// The duration of the eligible role assignment in hours.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.113.0/docs/resources/pim_eligible_role_assignment#duration_hours PimEligibleRoleAssignment#duration_hours}
	DurationHours *float64 `field:"optional" json:"durationHours" yaml:"durationHours"`
	// The end date/time of the eligible role assignment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.113.0/docs/resources/pim_eligible_role_assignment#end_date_time PimEligibleRoleAssignment#end_date_time}
	EndDateTime *string `field:"optional" json:"endDateTime" yaml:"endDateTime"`
}

