package cdnfrontdoorsecuritypolicy


type CdnFrontdoorSecurityPolicyTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_frontdoor_security_policy#create CdnFrontdoorSecurityPolicy#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_frontdoor_security_policy#delete CdnFrontdoorSecurityPolicy#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_frontdoor_security_policy#read CdnFrontdoorSecurityPolicy#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}

