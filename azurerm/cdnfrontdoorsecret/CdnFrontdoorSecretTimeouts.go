package cdnfrontdoorsecret


type CdnFrontdoorSecretTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_frontdoor_secret#create CdnFrontdoorSecret#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_frontdoor_secret#delete CdnFrontdoorSecret#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_frontdoor_secret#read CdnFrontdoorSecret#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}

