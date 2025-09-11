// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package monitoralertprocessingruleactiongroup


type MonitorAlertProcessingRuleActionGroupSchedule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs/resources/monitor_alert_processing_rule_action_group#effective_from MonitorAlertProcessingRuleActionGroup#effective_from}.
	EffectiveFrom *string `field:"optional" json:"effectiveFrom" yaml:"effectiveFrom"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs/resources/monitor_alert_processing_rule_action_group#effective_until MonitorAlertProcessingRuleActionGroup#effective_until}.
	EffectiveUntil *string `field:"optional" json:"effectiveUntil" yaml:"effectiveUntil"`
	// recurrence block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs/resources/monitor_alert_processing_rule_action_group#recurrence MonitorAlertProcessingRuleActionGroup#recurrence}
	Recurrence *MonitorAlertProcessingRuleActionGroupScheduleRecurrence `field:"optional" json:"recurrence" yaml:"recurrence"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs/resources/monitor_alert_processing_rule_action_group#time_zone MonitorAlertProcessingRuleActionGroup#time_zone}.
	TimeZone *string `field:"optional" json:"timeZone" yaml:"timeZone"`
}

