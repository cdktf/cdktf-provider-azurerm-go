// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sentinelautomationrule


type SentinelAutomationRuleActionIncidentTask struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/sentinel_automation_rule#order SentinelAutomationRule#order}.
	Order *float64 `field:"required" json:"order" yaml:"order"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/sentinel_automation_rule#title SentinelAutomationRule#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/sentinel_automation_rule#description SentinelAutomationRule#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

