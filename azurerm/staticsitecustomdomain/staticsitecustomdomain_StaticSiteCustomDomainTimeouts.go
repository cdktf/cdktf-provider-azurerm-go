package staticsitecustomdomain


type StaticSiteCustomDomainTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/static_site_custom_domain#create StaticSiteCustomDomain#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/static_site_custom_domain#delete StaticSiteCustomDomain#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/static_site_custom_domain#read StaticSiteCustomDomain#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/static_site_custom_domain#update StaticSiteCustomDomain#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}
