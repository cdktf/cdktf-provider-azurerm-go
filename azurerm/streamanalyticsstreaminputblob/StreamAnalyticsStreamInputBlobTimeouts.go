package streamanalyticsstreaminputblob


type StreamAnalyticsStreamInputBlobTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/stream_analytics_stream_input_blob#create StreamAnalyticsStreamInputBlob#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/stream_analytics_stream_input_blob#delete StreamAnalyticsStreamInputBlob#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/stream_analytics_stream_input_blob#read StreamAnalyticsStreamInputBlob#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/stream_analytics_stream_input_blob#update StreamAnalyticsStreamInputBlob#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}
