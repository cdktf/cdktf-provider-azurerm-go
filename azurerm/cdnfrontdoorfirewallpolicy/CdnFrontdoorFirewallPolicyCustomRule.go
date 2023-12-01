// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cdnfrontdoorfirewallpolicy


type CdnFrontdoorFirewallPolicyCustomRule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.83.0/docs/resources/cdn_frontdoor_firewall_policy#action CdnFrontdoorFirewallPolicy#action}.
	Action *string `field:"required" json:"action" yaml:"action"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.83.0/docs/resources/cdn_frontdoor_firewall_policy#name CdnFrontdoorFirewallPolicy#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.83.0/docs/resources/cdn_frontdoor_firewall_policy#type CdnFrontdoorFirewallPolicy#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.83.0/docs/resources/cdn_frontdoor_firewall_policy#enabled CdnFrontdoorFirewallPolicy#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// match_condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.83.0/docs/resources/cdn_frontdoor_firewall_policy#match_condition CdnFrontdoorFirewallPolicy#match_condition}
	MatchCondition interface{} `field:"optional" json:"matchCondition" yaml:"matchCondition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.83.0/docs/resources/cdn_frontdoor_firewall_policy#priority CdnFrontdoorFirewallPolicy#priority}.
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.83.0/docs/resources/cdn_frontdoor_firewall_policy#rate_limit_duration_in_minutes CdnFrontdoorFirewallPolicy#rate_limit_duration_in_minutes}.
	RateLimitDurationInMinutes *float64 `field:"optional" json:"rateLimitDurationInMinutes" yaml:"rateLimitDurationInMinutes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.83.0/docs/resources/cdn_frontdoor_firewall_policy#rate_limit_threshold CdnFrontdoorFirewallPolicy#rate_limit_threshold}.
	RateLimitThreshold *float64 `field:"optional" json:"rateLimitThreshold" yaml:"rateLimitThreshold"`
}

