// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package webapplicationfirewallpolicy


type WebApplicationFirewallPolicyManagedRulesManagedRuleSetRuleGroupOverride struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.31.0/docs/resources/web_application_firewall_policy#rule_group_name WebApplicationFirewallPolicy#rule_group_name}.
	RuleGroupName *string `field:"required" json:"ruleGroupName" yaml:"ruleGroupName"`
	// rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.31.0/docs/resources/web_application_firewall_policy#rule WebApplicationFirewallPolicy#rule}
	Rule interface{} `field:"optional" json:"rule" yaml:"rule"`
}

