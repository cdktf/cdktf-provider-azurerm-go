package monitordatacollectionrule


type MonitorDataCollectionRuleDataSourcesDataImport struct {
	// event_hub_data_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.56.0/docs/resources/monitor_data_collection_rule#event_hub_data_source MonitorDataCollectionRule#event_hub_data_source}
	EventHubDataSource interface{} `field:"required" json:"eventHubDataSource" yaml:"eventHubDataSource"`
}

