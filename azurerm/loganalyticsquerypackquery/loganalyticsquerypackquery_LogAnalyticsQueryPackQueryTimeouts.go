package loganalyticsquerypackquery


type LogAnalyticsQueryPackQueryTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/log_analytics_query_pack_query#create LogAnalyticsQueryPackQuery#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/log_analytics_query_pack_query#delete LogAnalyticsQueryPackQuery#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/log_analytics_query_pack_query#read LogAnalyticsQueryPackQuery#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/log_analytics_query_pack_query#update LogAnalyticsQueryPackQuery#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}
