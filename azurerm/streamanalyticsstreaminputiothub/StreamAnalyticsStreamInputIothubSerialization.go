package streamanalyticsstreaminputiothub


type StreamAnalyticsStreamInputIothubSerialization struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.56.0/docs/resources/stream_analytics_stream_input_iothub#type StreamAnalyticsStreamInputIothub#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.56.0/docs/resources/stream_analytics_stream_input_iothub#encoding StreamAnalyticsStreamInputIothub#encoding}.
	Encoding *string `field:"optional" json:"encoding" yaml:"encoding"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.56.0/docs/resources/stream_analytics_stream_input_iothub#field_delimiter StreamAnalyticsStreamInputIothub#field_delimiter}.
	FieldDelimiter *string `field:"optional" json:"fieldDelimiter" yaml:"fieldDelimiter"`
}

