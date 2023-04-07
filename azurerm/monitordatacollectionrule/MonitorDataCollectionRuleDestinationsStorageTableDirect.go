package monitordatacollectionrule


type MonitorDataCollectionRuleDestinationsStorageTableDirect struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_data_collection_rule#name MonitorDataCollectionRule#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_data_collection_rule#storage_account_id MonitorDataCollectionRule#storage_account_id}.
	StorageAccountId *string `field:"required" json:"storageAccountId" yaml:"storageAccountId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_data_collection_rule#table_name MonitorDataCollectionRule#table_name}.
	TableName *string `field:"required" json:"tableName" yaml:"tableName"`
}

