package mediastreamingendpoint


type MediaStreamingEndpointTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/media_streaming_endpoint#create MediaStreamingEndpoint#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/media_streaming_endpoint#delete MediaStreamingEndpoint#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/media_streaming_endpoint#read MediaStreamingEndpoint#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/media_streaming_endpoint#update MediaStreamingEndpoint#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}
