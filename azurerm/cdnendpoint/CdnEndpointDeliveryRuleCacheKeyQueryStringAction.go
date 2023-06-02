package cdnendpoint


type CdnEndpointDeliveryRuleCacheKeyQueryStringAction struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.59.0/docs/resources/cdn_endpoint#behavior CdnEndpoint#behavior}.
	Behavior *string `field:"required" json:"behavior" yaml:"behavior"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.59.0/docs/resources/cdn_endpoint#parameters CdnEndpoint#parameters}.
	Parameters *string `field:"optional" json:"parameters" yaml:"parameters"`
}

