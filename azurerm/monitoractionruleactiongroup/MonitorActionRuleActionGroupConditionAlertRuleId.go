// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package monitoractionruleactiongroup


type MonitorActionRuleActionGroupConditionAlertRuleId struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.102.0/docs/resources/monitor_action_rule_action_group#operator MonitorActionRuleActionGroup#operator}.
	Operator *string `field:"required" json:"operator" yaml:"operator"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.102.0/docs/resources/monitor_action_rule_action_group#values MonitorActionRuleActionGroup#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

