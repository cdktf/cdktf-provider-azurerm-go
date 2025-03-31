// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cdnfrontdoorsecuritypolicy


type CdnFrontdoorSecurityPolicySecurityPolicies struct {
	// firewall block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.25.0/docs/resources/cdn_frontdoor_security_policy#firewall CdnFrontdoorSecurityPolicy#firewall}
	Firewall *CdnFrontdoorSecurityPolicySecurityPoliciesFirewall `field:"required" json:"firewall" yaml:"firewall"`
}

