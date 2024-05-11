// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sentinelautomationrule


type SentinelAutomationRuleCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/sentinel_automation_rule#operator SentinelAutomationRule#operator}.
	Operator *string `field:"required" json:"operator" yaml:"operator"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/sentinel_automation_rule#property SentinelAutomationRule#property}.
	Property *string `field:"required" json:"property" yaml:"property"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/sentinel_automation_rule#values SentinelAutomationRule#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

