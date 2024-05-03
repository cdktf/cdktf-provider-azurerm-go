// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package securitycenterautomation


type SecurityCenterAutomationSource struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.102.0/docs/resources/security_center_automation#event_source SecurityCenterAutomation#event_source}.
	EventSource *string `field:"required" json:"eventSource" yaml:"eventSource"`
	// rule_set block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.102.0/docs/resources/security_center_automation#rule_set SecurityCenterAutomation#rule_set}
	RuleSet interface{} `field:"optional" json:"ruleSet" yaml:"ruleSet"`
}

