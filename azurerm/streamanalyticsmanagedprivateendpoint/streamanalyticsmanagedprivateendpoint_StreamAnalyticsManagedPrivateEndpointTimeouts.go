package streamanalyticsmanagedprivateendpoint


type StreamAnalyticsManagedPrivateEndpointTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/stream_analytics_managed_private_endpoint#create StreamAnalyticsManagedPrivateEndpoint#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/stream_analytics_managed_private_endpoint#delete StreamAnalyticsManagedPrivateEndpoint#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/stream_analytics_managed_private_endpoint#read StreamAnalyticsManagedPrivateEndpoint#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}
