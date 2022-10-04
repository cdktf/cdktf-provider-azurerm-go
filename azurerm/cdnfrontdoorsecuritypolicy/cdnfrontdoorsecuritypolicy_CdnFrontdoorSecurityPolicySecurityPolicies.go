package cdnfrontdoorsecuritypolicy


type CdnFrontdoorSecurityPolicySecurityPolicies struct {
	// firewall block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_frontdoor_security_policy#firewall CdnFrontdoorSecurityPolicy#firewall}
	Firewall *CdnFrontdoorSecurityPolicySecurityPoliciesFirewall `field:"required" json:"firewall" yaml:"firewall"`
}

