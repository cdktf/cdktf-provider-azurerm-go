package monitordatacollectionruleassociation


type MonitorDataCollectionRuleAssociationTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_data_collection_rule_association#create MonitorDataCollectionRuleAssociation#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_data_collection_rule_association#delete MonitorDataCollectionRuleAssociation#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_data_collection_rule_association#read MonitorDataCollectionRuleAssociation#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_data_collection_rule_association#update MonitorDataCollectionRuleAssociation#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}
