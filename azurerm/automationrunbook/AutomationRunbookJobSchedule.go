// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package automationrunbook


type AutomationRunbookJobSchedule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.22.0/docs/resources/automation_runbook#job_schedule_id AutomationRunbook#job_schedule_id}.
	JobScheduleId *string `field:"optional" json:"jobScheduleId" yaml:"jobScheduleId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.22.0/docs/resources/automation_runbook#parameters AutomationRunbook#parameters}.
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.22.0/docs/resources/automation_runbook#run_on AutomationRunbook#run_on}.
	RunOn *string `field:"optional" json:"runOn" yaml:"runOn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.22.0/docs/resources/automation_runbook#schedule_name AutomationRunbook#schedule_name}.
	ScheduleName *string `field:"optional" json:"scheduleName" yaml:"scheduleName"`
}

