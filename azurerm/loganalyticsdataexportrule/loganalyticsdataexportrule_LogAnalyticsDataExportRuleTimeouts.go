package loganalyticsdataexportrule


type LogAnalyticsDataExportRuleTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/log_analytics_data_export_rule#create LogAnalyticsDataExportRule#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/log_analytics_data_export_rule#delete LogAnalyticsDataExportRule#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/log_analytics_data_export_rule#read LogAnalyticsDataExportRule#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/log_analytics_data_export_rule#update LogAnalyticsDataExportRule#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}
