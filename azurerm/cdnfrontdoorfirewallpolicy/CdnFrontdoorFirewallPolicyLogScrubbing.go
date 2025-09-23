// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cdnfrontdoorfirewallpolicy


type CdnFrontdoorFirewallPolicyLogScrubbing struct {
	// scrubbing_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/cdn_frontdoor_firewall_policy#scrubbing_rule CdnFrontdoorFirewallPolicy#scrubbing_rule}
	ScrubbingRule interface{} `field:"required" json:"scrubbingRule" yaml:"scrubbingRule"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/cdn_frontdoor_firewall_policy#enabled CdnFrontdoorFirewallPolicy#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

