package loganalyticsquerypack


type LogAnalyticsQueryPackTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/log_analytics_query_pack#create LogAnalyticsQueryPack#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/log_analytics_query_pack#delete LogAnalyticsQueryPack#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/log_analytics_query_pack#read LogAnalyticsQueryPack#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/log_analytics_query_pack#update LogAnalyticsQueryPack#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}
