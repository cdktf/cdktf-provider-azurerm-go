package streamanalyticscluster


type StreamAnalyticsClusterTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/stream_analytics_cluster#create StreamAnalyticsCluster#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/stream_analytics_cluster#delete StreamAnalyticsCluster#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/stream_analytics_cluster#read StreamAnalyticsCluster#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/stream_analytics_cluster#update StreamAnalyticsCluster#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}
