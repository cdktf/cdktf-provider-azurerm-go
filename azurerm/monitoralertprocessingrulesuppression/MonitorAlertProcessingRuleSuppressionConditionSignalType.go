// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package monitoralertprocessingrulesuppression


type MonitorAlertProcessingRuleSuppressionConditionSignalType struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.30.0/docs/resources/monitor_alert_processing_rule_suppression#operator MonitorAlertProcessingRuleSuppression#operator}.
	Operator *string `field:"required" json:"operator" yaml:"operator"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.30.0/docs/resources/monitor_alert_processing_rule_suppression#values MonitorAlertProcessingRuleSuppression#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

