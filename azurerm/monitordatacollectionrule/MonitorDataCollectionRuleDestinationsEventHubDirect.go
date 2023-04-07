package monitordatacollectionrule


type MonitorDataCollectionRuleDestinationsEventHubDirect struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_data_collection_rule#event_hub_id MonitorDataCollectionRule#event_hub_id}.
	EventHubId *string `field:"required" json:"eventHubId" yaml:"eventHubId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_data_collection_rule#name MonitorDataCollectionRule#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
}

