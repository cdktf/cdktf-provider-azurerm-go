package streamanalyticsoutputcosmosdb


type StreamAnalyticsOutputCosmosdbTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/stream_analytics_output_cosmosdb#create StreamAnalyticsOutputCosmosdb#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/stream_analytics_output_cosmosdb#delete StreamAnalyticsOutputCosmosdb#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/stream_analytics_output_cosmosdb#read StreamAnalyticsOutputCosmosdb#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/stream_analytics_output_cosmosdb#update StreamAnalyticsOutputCosmosdb#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}
