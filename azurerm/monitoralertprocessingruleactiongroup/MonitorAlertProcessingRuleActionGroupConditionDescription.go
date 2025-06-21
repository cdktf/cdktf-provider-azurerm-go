// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package monitoralertprocessingruleactiongroup


type MonitorAlertProcessingRuleActionGroupConditionDescription struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/monitor_alert_processing_rule_action_group#operator MonitorAlertProcessingRuleActionGroup#operator}.
	Operator *string `field:"required" json:"operator" yaml:"operator"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/monitor_alert_processing_rule_action_group#values MonitorAlertProcessingRuleActionGroup#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

