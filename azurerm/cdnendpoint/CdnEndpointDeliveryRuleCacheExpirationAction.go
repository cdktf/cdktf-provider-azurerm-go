package cdnendpoint


type CdnEndpointDeliveryRuleCacheExpirationAction struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.68.0/docs/resources/cdn_endpoint#behavior CdnEndpoint#behavior}.
	Behavior *string `field:"required" json:"behavior" yaml:"behavior"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.68.0/docs/resources/cdn_endpoint#duration CdnEndpoint#duration}.
	Duration *string `field:"optional" json:"duration" yaml:"duration"`
}

