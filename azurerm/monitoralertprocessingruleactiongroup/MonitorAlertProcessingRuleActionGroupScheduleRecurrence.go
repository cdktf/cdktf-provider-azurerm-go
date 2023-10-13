// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package monitoralertprocessingruleactiongroup


type MonitorAlertProcessingRuleActionGroupScheduleRecurrence struct {
	// daily block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/resources/monitor_alert_processing_rule_action_group#daily MonitorAlertProcessingRuleActionGroup#daily}
	Daily interface{} `field:"optional" json:"daily" yaml:"daily"`
	// monthly block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/resources/monitor_alert_processing_rule_action_group#monthly MonitorAlertProcessingRuleActionGroup#monthly}
	Monthly interface{} `field:"optional" json:"monthly" yaml:"monthly"`
	// weekly block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/resources/monitor_alert_processing_rule_action_group#weekly MonitorAlertProcessingRuleActionGroup#weekly}
	Weekly interface{} `field:"optional" json:"weekly" yaml:"weekly"`
}

