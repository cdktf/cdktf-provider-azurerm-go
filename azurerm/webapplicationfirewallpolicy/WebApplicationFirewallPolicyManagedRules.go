// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package webapplicationfirewallpolicy


type WebApplicationFirewallPolicyManagedRules struct {
	// managed_rule_set block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.2.0/docs/resources/web_application_firewall_policy#managed_rule_set WebApplicationFirewallPolicy#managed_rule_set}
	ManagedRuleSet interface{} `field:"required" json:"managedRuleSet" yaml:"managedRuleSet"`
	// exclusion block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.2.0/docs/resources/web_application_firewall_policy#exclusion WebApplicationFirewallPolicy#exclusion}
	Exclusion interface{} `field:"optional" json:"exclusion" yaml:"exclusion"`
}

