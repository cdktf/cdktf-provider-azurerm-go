package monitordatacollectionrule


type MonitorDataCollectionRuleDataSourcesLogFile struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_data_collection_rule#file_patterns MonitorDataCollectionRule#file_patterns}.
	FilePatterns *[]*string `field:"required" json:"filePatterns" yaml:"filePatterns"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_data_collection_rule#format MonitorDataCollectionRule#format}.
	Format *string `field:"required" json:"format" yaml:"format"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_data_collection_rule#name MonitorDataCollectionRule#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_data_collection_rule#streams MonitorDataCollectionRule#streams}.
	Streams *[]*string `field:"required" json:"streams" yaml:"streams"`
	// settings block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_data_collection_rule#settings MonitorDataCollectionRule#settings}
	Settings *MonitorDataCollectionRuleDataSourcesLogFileSettings `field:"optional" json:"settings" yaml:"settings"`
}

