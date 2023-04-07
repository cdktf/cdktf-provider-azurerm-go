package monitordatacollectionrule


type MonitorDataCollectionRuleDestinationsMonitorAccount struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_data_collection_rule#monitor_account_id MonitorDataCollectionRule#monitor_account_id}.
	MonitorAccountId *string `field:"required" json:"monitorAccountId" yaml:"monitorAccountId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_data_collection_rule#name MonitorDataCollectionRule#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
}

