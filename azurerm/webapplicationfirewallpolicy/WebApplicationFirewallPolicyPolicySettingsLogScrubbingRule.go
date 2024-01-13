// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package webapplicationfirewallpolicy


type WebApplicationFirewallPolicyPolicySettingsLogScrubbingRule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.87.0/docs/resources/web_application_firewall_policy#match_variable WebApplicationFirewallPolicy#match_variable}.
	MatchVariable *string `field:"required" json:"matchVariable" yaml:"matchVariable"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.87.0/docs/resources/web_application_firewall_policy#enabled WebApplicationFirewallPolicy#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// When matchVariable is a collection, operator used to specify which elements in the collection this rule applies to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.87.0/docs/resources/web_application_firewall_policy#selector WebApplicationFirewallPolicy#selector}
	Selector *string `field:"optional" json:"selector" yaml:"selector"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.87.0/docs/resources/web_application_firewall_policy#selector_match_operator WebApplicationFirewallPolicy#selector_match_operator}.
	SelectorMatchOperator *string `field:"optional" json:"selectorMatchOperator" yaml:"selectorMatchOperator"`
}

