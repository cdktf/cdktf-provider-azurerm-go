// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package automationschedule


type AutomationScheduleMonthlyOccurrence struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/automation_schedule#day AutomationSchedule#day}.
	Day *string `field:"required" json:"day" yaml:"day"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/automation_schedule#occurrence AutomationSchedule#occurrence}.
	Occurrence *float64 `field:"required" json:"occurrence" yaml:"occurrence"`
}

