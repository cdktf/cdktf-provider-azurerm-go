// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pimactiveroleassignment


type PimActiveRoleAssignmentScheduleExpiration struct {
	// The duration of the assignment in days.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/pim_active_role_assignment#duration_days PimActiveRoleAssignment#duration_days}
	DurationDays *float64 `field:"optional" json:"durationDays" yaml:"durationDays"`
	// The duration of the assignment in hours.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/pim_active_role_assignment#duration_hours PimActiveRoleAssignment#duration_hours}
	DurationHours *float64 `field:"optional" json:"durationHours" yaml:"durationHours"`
	// The end date time of the assignment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/pim_active_role_assignment#end_date_time PimActiveRoleAssignment#end_date_time}
	EndDateTime *string `field:"optional" json:"endDateTime" yaml:"endDateTime"`
}

