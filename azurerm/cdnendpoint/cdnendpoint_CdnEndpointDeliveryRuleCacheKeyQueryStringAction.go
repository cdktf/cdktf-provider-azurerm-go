package cdnendpoint


type CdnEndpointDeliveryRuleCacheKeyQueryStringAction struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#behavior CdnEndpoint#behavior}.
	Behavior *string `field:"required" json:"behavior" yaml:"behavior"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#parameters CdnEndpoint#parameters}.
	Parameters *string `field:"optional" json:"parameters" yaml:"parameters"`
}
