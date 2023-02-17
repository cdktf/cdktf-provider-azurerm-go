package cdnendpoint


type CdnEndpointGeoFilter struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#action CdnEndpoint#action}.
	Action *string `field:"required" json:"action" yaml:"action"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#country_codes CdnEndpoint#country_codes}.
	CountryCodes *[]*string `field:"required" json:"countryCodes" yaml:"countryCodes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#relative_path CdnEndpoint#relative_path}.
	RelativePath *string `field:"required" json:"relativePath" yaml:"relativePath"`
}

