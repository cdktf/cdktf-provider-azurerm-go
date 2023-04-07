package monitordatacollectionrule


type MonitorDataCollectionRuleStreamDeclarationColumn struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_data_collection_rule#name MonitorDataCollectionRule#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_data_collection_rule#type MonitorDataCollectionRule#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}

