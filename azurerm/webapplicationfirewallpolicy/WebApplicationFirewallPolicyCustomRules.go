// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package webapplicationfirewallpolicy


type WebApplicationFirewallPolicyCustomRules struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.47.0/docs/resources/web_application_firewall_policy#action WebApplicationFirewallPolicy#action}.
	Action *string `field:"required" json:"action" yaml:"action"`
	// match_conditions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.47.0/docs/resources/web_application_firewall_policy#match_conditions WebApplicationFirewallPolicy#match_conditions}
	MatchConditions interface{} `field:"required" json:"matchConditions" yaml:"matchConditions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.47.0/docs/resources/web_application_firewall_policy#priority WebApplicationFirewallPolicy#priority}.
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.47.0/docs/resources/web_application_firewall_policy#rule_type WebApplicationFirewallPolicy#rule_type}.
	RuleType *string `field:"required" json:"ruleType" yaml:"ruleType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.47.0/docs/resources/web_application_firewall_policy#enabled WebApplicationFirewallPolicy#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.47.0/docs/resources/web_application_firewall_policy#group_rate_limit_by WebApplicationFirewallPolicy#group_rate_limit_by}.
	GroupRateLimitBy *string `field:"optional" json:"groupRateLimitBy" yaml:"groupRateLimitBy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.47.0/docs/resources/web_application_firewall_policy#name WebApplicationFirewallPolicy#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.47.0/docs/resources/web_application_firewall_policy#rate_limit_duration WebApplicationFirewallPolicy#rate_limit_duration}.
	RateLimitDuration *string `field:"optional" json:"rateLimitDuration" yaml:"rateLimitDuration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.47.0/docs/resources/web_application_firewall_policy#rate_limit_threshold WebApplicationFirewallPolicy#rate_limit_threshold}.
	RateLimitThreshold *float64 `field:"optional" json:"rateLimitThreshold" yaml:"rateLimitThreshold"`
}

