// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package webapplicationfirewallpolicy


type WebApplicationFirewallPolicyManagedRulesExclusion struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.15.0/docs/resources/web_application_firewall_policy#match_variable WebApplicationFirewallPolicy#match_variable}.
	MatchVariable *string `field:"required" json:"matchVariable" yaml:"matchVariable"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.15.0/docs/resources/web_application_firewall_policy#selector WebApplicationFirewallPolicy#selector}.
	Selector *string `field:"required" json:"selector" yaml:"selector"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.15.0/docs/resources/web_application_firewall_policy#selector_match_operator WebApplicationFirewallPolicy#selector_match_operator}.
	SelectorMatchOperator *string `field:"required" json:"selectorMatchOperator" yaml:"selectorMatchOperator"`
	// excluded_rule_set block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.15.0/docs/resources/web_application_firewall_policy#excluded_rule_set WebApplicationFirewallPolicy#excluded_rule_set}
	ExcludedRuleSet *WebApplicationFirewallPolicyManagedRulesExclusionExcludedRuleSet `field:"optional" json:"excludedRuleSet" yaml:"excludedRuleSet"`
}

