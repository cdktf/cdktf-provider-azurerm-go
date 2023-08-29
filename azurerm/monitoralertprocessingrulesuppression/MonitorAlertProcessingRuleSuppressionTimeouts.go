// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package monitoralertprocessingrulesuppression


type MonitorAlertProcessingRuleSuppressionTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/monitor_alert_processing_rule_suppression#create MonitorAlertProcessingRuleSuppression#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/monitor_alert_processing_rule_suppression#delete MonitorAlertProcessingRuleSuppression#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/monitor_alert_processing_rule_suppression#read MonitorAlertProcessingRuleSuppression#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/monitor_alert_processing_rule_suppression#update MonitorAlertProcessingRuleSuppression#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

