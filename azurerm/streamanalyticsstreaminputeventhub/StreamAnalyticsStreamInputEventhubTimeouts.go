package streamanalyticsstreaminputeventhub


type StreamAnalyticsStreamInputEventhubTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/stream_analytics_stream_input_eventhub#create StreamAnalyticsStreamInputEventhub#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/stream_analytics_stream_input_eventhub#delete StreamAnalyticsStreamInputEventhub#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/stream_analytics_stream_input_eventhub#read StreamAnalyticsStreamInputEventhub#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/stream_analytics_stream_input_eventhub#update StreamAnalyticsStreamInputEventhub#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}
