package monitordatacollectionrule


type MonitorDataCollectionRuleDataSources struct {
	// extension block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_data_collection_rule#extension MonitorDataCollectionRule#extension}
	Extension interface{} `field:"optional" json:"extension" yaml:"extension"`
	// performance_counter block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_data_collection_rule#performance_counter MonitorDataCollectionRule#performance_counter}
	PerformanceCounter interface{} `field:"optional" json:"performanceCounter" yaml:"performanceCounter"`
	// syslog block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_data_collection_rule#syslog MonitorDataCollectionRule#syslog}
	Syslog interface{} `field:"optional" json:"syslog" yaml:"syslog"`
	// windows_event_log block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_data_collection_rule#windows_event_log MonitorDataCollectionRule#windows_event_log}
	WindowsEventLog interface{} `field:"optional" json:"windowsEventLog" yaml:"windowsEventLog"`
}

