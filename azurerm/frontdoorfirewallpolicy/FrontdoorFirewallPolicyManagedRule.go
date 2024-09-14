// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package frontdoorfirewallpolicy


type FrontdoorFirewallPolicyManagedRule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.2.0/docs/resources/frontdoor_firewall_policy#type FrontdoorFirewallPolicy#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.2.0/docs/resources/frontdoor_firewall_policy#version FrontdoorFirewallPolicy#version}.
	Version *string `field:"required" json:"version" yaml:"version"`
	// exclusion block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.2.0/docs/resources/frontdoor_firewall_policy#exclusion FrontdoorFirewallPolicy#exclusion}
	Exclusion interface{} `field:"optional" json:"exclusion" yaml:"exclusion"`
	// override block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.2.0/docs/resources/frontdoor_firewall_policy#override FrontdoorFirewallPolicy#override}
	Override interface{} `field:"optional" json:"override" yaml:"override"`
}

