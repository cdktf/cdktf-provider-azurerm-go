package cdnendpoint


type CdnEndpointDeliveryRuleCacheExpirationAction struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#behavior CdnEndpoint#behavior}.
	Behavior *string `field:"required" json:"behavior" yaml:"behavior"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#duration CdnEndpoint#duration}.
	Duration *string `field:"optional" json:"duration" yaml:"duration"`
}

