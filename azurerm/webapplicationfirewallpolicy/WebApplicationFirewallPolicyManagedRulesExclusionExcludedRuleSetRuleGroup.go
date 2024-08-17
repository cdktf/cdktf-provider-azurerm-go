// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package webapplicationfirewallpolicy


type WebApplicationFirewallPolicyManagedRulesExclusionExcludedRuleSetRuleGroup struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/web_application_firewall_policy#rule_group_name WebApplicationFirewallPolicy#rule_group_name}.
	RuleGroupName *string `field:"required" json:"ruleGroupName" yaml:"ruleGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/web_application_firewall_policy#excluded_rules WebApplicationFirewallPolicy#excluded_rules}.
	ExcludedRules *[]*string `field:"optional" json:"excludedRules" yaml:"excludedRules"`
}

