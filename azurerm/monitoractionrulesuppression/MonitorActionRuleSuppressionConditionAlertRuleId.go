// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package monitoractionrulesuppression


type MonitorActionRuleSuppressionConditionAlertRuleId struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.88.0/docs/resources/monitor_action_rule_suppression#operator MonitorActionRuleSuppression#operator}.
	Operator *string `field:"required" json:"operator" yaml:"operator"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.88.0/docs/resources/monitor_action_rule_suppression#values MonitorActionRuleSuppression#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

