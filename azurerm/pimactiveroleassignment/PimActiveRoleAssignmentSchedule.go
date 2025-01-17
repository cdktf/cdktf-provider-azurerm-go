// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pimactiveroleassignment


type PimActiveRoleAssignmentSchedule struct {
	// expiration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.16.0/docs/resources/pim_active_role_assignment#expiration PimActiveRoleAssignment#expiration}
	Expiration *PimActiveRoleAssignmentScheduleExpiration `field:"optional" json:"expiration" yaml:"expiration"`
	// The start date/time of the role assignment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.16.0/docs/resources/pim_active_role_assignment#start_date_time PimActiveRoleAssignment#start_date_time}
	StartDateTime *string `field:"optional" json:"startDateTime" yaml:"startDateTime"`
}

