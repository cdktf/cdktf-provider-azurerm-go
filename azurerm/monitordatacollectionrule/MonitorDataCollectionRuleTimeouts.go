package monitordatacollectionrule


type MonitorDataCollectionRuleTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_data_collection_rule#create MonitorDataCollectionRule#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_data_collection_rule#delete MonitorDataCollectionRule#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_data_collection_rule#read MonitorDataCollectionRule#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_data_collection_rule#update MonitorDataCollectionRule#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}
