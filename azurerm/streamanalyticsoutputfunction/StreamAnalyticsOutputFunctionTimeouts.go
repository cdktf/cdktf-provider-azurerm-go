package streamanalyticsoutputfunction


type StreamAnalyticsOutputFunctionTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/stream_analytics_output_function#create StreamAnalyticsOutputFunction#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/stream_analytics_output_function#delete StreamAnalyticsOutputFunction#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/stream_analytics_output_function#read StreamAnalyticsOutputFunction#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/stream_analytics_output_function#update StreamAnalyticsOutputFunction#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}
