package monitordatacollectionrule


type MonitorDataCollectionRuleDataSourcesPlatformTelemetry struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/monitor_data_collection_rule#name MonitorDataCollectionRule#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/monitor_data_collection_rule#streams MonitorDataCollectionRule#streams}.
	Streams *[]*string `field:"required" json:"streams" yaml:"streams"`
}

