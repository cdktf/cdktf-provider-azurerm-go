// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cdnfrontdoorsecuritypolicy


type CdnFrontdoorSecurityPolicySecurityPoliciesFirewall struct {
	// association block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.33.0/docs/resources/cdn_frontdoor_security_policy#association CdnFrontdoorSecurityPolicy#association}
	Association *CdnFrontdoorSecurityPolicySecurityPoliciesFirewallAssociation `field:"required" json:"association" yaml:"association"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.33.0/docs/resources/cdn_frontdoor_security_policy#cdn_frontdoor_firewall_policy_id CdnFrontdoorSecurityPolicy#cdn_frontdoor_firewall_policy_id}.
	CdnFrontdoorFirewallPolicyId *string `field:"required" json:"cdnFrontdoorFirewallPolicyId" yaml:"cdnFrontdoorFirewallPolicyId"`
}

