// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package monitoralertprocessingruleactiongroup


type MonitorAlertProcessingRuleActionGroupScheduleRecurrenceWeekly struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.79.0/docs/resources/monitor_alert_processing_rule_action_group#days_of_week MonitorAlertProcessingRuleActionGroup#days_of_week}.
	DaysOfWeek *[]*string `field:"required" json:"daysOfWeek" yaml:"daysOfWeek"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.79.0/docs/resources/monitor_alert_processing_rule_action_group#end_time MonitorAlertProcessingRuleActionGroup#end_time}.
	EndTime *string `field:"optional" json:"endTime" yaml:"endTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.79.0/docs/resources/monitor_alert_processing_rule_action_group#start_time MonitorAlertProcessingRuleActionGroup#start_time}.
	StartTime *string `field:"optional" json:"startTime" yaml:"startTime"`
}

