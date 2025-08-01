// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package monitoralertprocessingruleactiongroup


type MonitorAlertProcessingRuleActionGroupScheduleRecurrenceDaily struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/monitor_alert_processing_rule_action_group#end_time MonitorAlertProcessingRuleActionGroup#end_time}.
	EndTime *string `field:"required" json:"endTime" yaml:"endTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/monitor_alert_processing_rule_action_group#start_time MonitorAlertProcessingRuleActionGroup#start_time}.
	StartTime *string `field:"required" json:"startTime" yaml:"startTime"`
}

