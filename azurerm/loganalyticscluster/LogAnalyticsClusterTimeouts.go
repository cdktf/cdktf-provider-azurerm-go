package loganalyticscluster


type LogAnalyticsClusterTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/log_analytics_cluster#create LogAnalyticsCluster#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/log_analytics_cluster#delete LogAnalyticsCluster#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/log_analytics_cluster#read LogAnalyticsCluster#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/log_analytics_cluster#update LogAnalyticsCluster#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}
