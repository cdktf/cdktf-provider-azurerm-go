// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package webapplicationfirewallpolicy


type WebApplicationFirewallPolicyManagedRulesManagedRuleSet struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.95.0/docs/resources/web_application_firewall_policy#version WebApplicationFirewallPolicy#version}.
	Version *string `field:"required" json:"version" yaml:"version"`
	// rule_group_override block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.95.0/docs/resources/web_application_firewall_policy#rule_group_override WebApplicationFirewallPolicy#rule_group_override}
	RuleGroupOverride interface{} `field:"optional" json:"ruleGroupOverride" yaml:"ruleGroupOverride"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.95.0/docs/resources/web_application_firewall_policy#type WebApplicationFirewallPolicy#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

