package monitordatacollectionrule


type MonitorDataCollectionRuleDataSourcesDataImportEventHubDataSource struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.61.0/docs/resources/monitor_data_collection_rule#name MonitorDataCollectionRule#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.61.0/docs/resources/monitor_data_collection_rule#stream MonitorDataCollectionRule#stream}.
	Stream *string `field:"required" json:"stream" yaml:"stream"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.61.0/docs/resources/monitor_data_collection_rule#consumer_group MonitorDataCollectionRule#consumer_group}.
	ConsumerGroup *string `field:"optional" json:"consumerGroup" yaml:"consumerGroup"`
}

