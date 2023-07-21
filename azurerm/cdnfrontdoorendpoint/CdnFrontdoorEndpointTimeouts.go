package cdnfrontdoorendpoint


type CdnFrontdoorEndpointTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.66.0/docs/resources/cdn_frontdoor_endpoint#create CdnFrontdoorEndpoint#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.66.0/docs/resources/cdn_frontdoor_endpoint#delete CdnFrontdoorEndpoint#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.66.0/docs/resources/cdn_frontdoor_endpoint#read CdnFrontdoorEndpoint#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.66.0/docs/resources/cdn_frontdoor_endpoint#update CdnFrontdoorEndpoint#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

