package cdnfrontdoorsecuritypolicy


type CdnFrontdoorSecurityPolicySecurityPoliciesFirewall struct {
	// association block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_frontdoor_security_policy#association CdnFrontdoorSecurityPolicy#association}
	Association *CdnFrontdoorSecurityPolicySecurityPoliciesFirewallAssociation `field:"required" json:"association" yaml:"association"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_frontdoor_security_policy#cdn_frontdoor_firewall_policy_id CdnFrontdoorSecurityPolicy#cdn_frontdoor_firewall_policy_id}.
	CdnFrontdoorFirewallPolicyId *string `field:"required" json:"cdnFrontdoorFirewallPolicyId" yaml:"cdnFrontdoorFirewallPolicyId"`
}

