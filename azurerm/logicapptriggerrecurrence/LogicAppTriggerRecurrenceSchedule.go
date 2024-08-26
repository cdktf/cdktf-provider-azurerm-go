// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package logicapptriggerrecurrence


type LogicAppTriggerRecurrenceSchedule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.0.1/docs/resources/logic_app_trigger_recurrence#at_these_hours LogicAppTriggerRecurrence#at_these_hours}.
	AtTheseHours *[]*float64 `field:"optional" json:"atTheseHours" yaml:"atTheseHours"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.0.1/docs/resources/logic_app_trigger_recurrence#at_these_minutes LogicAppTriggerRecurrence#at_these_minutes}.
	AtTheseMinutes *[]*float64 `field:"optional" json:"atTheseMinutes" yaml:"atTheseMinutes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.0.1/docs/resources/logic_app_trigger_recurrence#on_these_days LogicAppTriggerRecurrence#on_these_days}.
	OnTheseDays *[]*string `field:"optional" json:"onTheseDays" yaml:"onTheseDays"`
}

