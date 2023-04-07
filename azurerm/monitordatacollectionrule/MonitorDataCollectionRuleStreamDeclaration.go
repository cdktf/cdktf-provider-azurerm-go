package monitordatacollectionrule


type MonitorDataCollectionRuleStreamDeclaration struct {
	// column block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_data_collection_rule#column MonitorDataCollectionRule#column}
	Column interface{} `field:"required" json:"column" yaml:"column"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_data_collection_rule#stream_name MonitorDataCollectionRule#stream_name}.
	StreamName *string `field:"required" json:"streamName" yaml:"streamName"`
}

