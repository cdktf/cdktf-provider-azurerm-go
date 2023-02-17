package cdnfrontdoorsecuritypolicy


type CdnFrontdoorSecurityPolicySecurityPoliciesFirewallAssociation struct {
	// domain block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_frontdoor_security_policy#domain CdnFrontdoorSecurityPolicy#domain}
	Domain interface{} `field:"required" json:"domain" yaml:"domain"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_frontdoor_security_policy#patterns_to_match CdnFrontdoorSecurityPolicy#patterns_to_match}.
	PatternsToMatch *[]*string `field:"required" json:"patternsToMatch" yaml:"patternsToMatch"`
}

