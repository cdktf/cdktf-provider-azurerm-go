// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package frontdoorfirewallpolicy


type FrontdoorFirewallPolicyManagedRuleOverrideExclusion struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.1/docs/resources/frontdoor_firewall_policy#match_variable FrontdoorFirewallPolicy#match_variable}.
	MatchVariable *string `field:"required" json:"matchVariable" yaml:"matchVariable"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.1/docs/resources/frontdoor_firewall_policy#operator FrontdoorFirewallPolicy#operator}.
	Operator *string `field:"required" json:"operator" yaml:"operator"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.1/docs/resources/frontdoor_firewall_policy#selector FrontdoorFirewallPolicy#selector}.
	Selector *string `field:"required" json:"selector" yaml:"selector"`
}

