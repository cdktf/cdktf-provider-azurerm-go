package streamanalyticsfunctionjavascriptudf


type StreamAnalyticsFunctionJavascriptUdfInput struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/stream_analytics_function_javascript_udf#type StreamAnalyticsFunctionJavascriptUdf#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/stream_analytics_function_javascript_udf#configuration_parameter StreamAnalyticsFunctionJavascriptUdf#configuration_parameter}.
	ConfigurationParameter interface{} `field:"optional" json:"configurationParameter" yaml:"configurationParameter"`
}

