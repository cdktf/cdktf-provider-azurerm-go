package monitordatacollectionrule


type MonitorDataCollectionRuleDataSourcesPrometheusForwarderLabelIncludeFilter struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_data_collection_rule#label MonitorDataCollectionRule#label}.
	Label *string `field:"required" json:"label" yaml:"label"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_data_collection_rule#value MonitorDataCollectionRule#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

