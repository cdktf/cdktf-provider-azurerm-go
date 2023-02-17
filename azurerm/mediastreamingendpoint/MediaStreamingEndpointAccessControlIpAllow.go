package mediastreamingendpoint


type MediaStreamingEndpointAccessControlIpAllow struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/media_streaming_endpoint#address MediaStreamingEndpoint#address}.
	Address *string `field:"optional" json:"address" yaml:"address"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/media_streaming_endpoint#name MediaStreamingEndpoint#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/media_streaming_endpoint#subnet_prefix_length MediaStreamingEndpoint#subnet_prefix_length}.
	SubnetPrefixLength *float64 `field:"optional" json:"subnetPrefixLength" yaml:"subnetPrefixLength"`
}

