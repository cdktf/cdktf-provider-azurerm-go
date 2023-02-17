package privatednsresolveroutboundendpoint


type PrivateDnsResolverOutboundEndpointTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/private_dns_resolver_outbound_endpoint#create PrivateDnsResolverOutboundEndpoint#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/private_dns_resolver_outbound_endpoint#delete PrivateDnsResolverOutboundEndpoint#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/private_dns_resolver_outbound_endpoint#read PrivateDnsResolverOutboundEndpoint#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/private_dns_resolver_outbound_endpoint#update PrivateDnsResolverOutboundEndpoint#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

