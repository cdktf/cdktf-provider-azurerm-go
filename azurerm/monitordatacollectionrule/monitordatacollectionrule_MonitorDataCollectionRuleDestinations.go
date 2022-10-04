package monitordatacollectionrule


type MonitorDataCollectionRuleDestinations struct {
	// azure_monitor_metrics block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_data_collection_rule#azure_monitor_metrics MonitorDataCollectionRule#azure_monitor_metrics}
	AzureMonitorMetrics *MonitorDataCollectionRuleDestinationsAzureMonitorMetrics `field:"optional" json:"azureMonitorMetrics" yaml:"azureMonitorMetrics"`
	// log_analytics block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_data_collection_rule#log_analytics MonitorDataCollectionRule#log_analytics}
	LogAnalytics interface{} `field:"optional" json:"logAnalytics" yaml:"logAnalytics"`
}

