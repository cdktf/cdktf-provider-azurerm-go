// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pimeligibleroleassignment


type PimEligibleRoleAssignmentSchedule struct {
	// expiration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.72.0/docs/resources/pim_eligible_role_assignment#expiration PimEligibleRoleAssignment#expiration}
	Expiration *PimEligibleRoleAssignmentScheduleExpiration `field:"optional" json:"expiration" yaml:"expiration"`
	// The start date time.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.72.0/docs/resources/pim_eligible_role_assignment#start_date_time PimEligibleRoleAssignment#start_date_time}
	StartDateTime *string `field:"optional" json:"startDateTime" yaml:"startDateTime"`
}

