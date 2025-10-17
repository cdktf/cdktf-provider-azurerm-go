// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package monitordatacollectionrule


type MonitorDataCollectionRuleDestinations struct {
	// azure_monitor_metrics block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.49.0/docs/resources/monitor_data_collection_rule#azure_monitor_metrics MonitorDataCollectionRule#azure_monitor_metrics}
	AzureMonitorMetrics *MonitorDataCollectionRuleDestinationsAzureMonitorMetrics `field:"optional" json:"azureMonitorMetrics" yaml:"azureMonitorMetrics"`
	// event_hub block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.49.0/docs/resources/monitor_data_collection_rule#event_hub MonitorDataCollectionRule#event_hub}
	EventHub *MonitorDataCollectionRuleDestinationsEventHub `field:"optional" json:"eventHub" yaml:"eventHub"`
	// event_hub_direct block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.49.0/docs/resources/monitor_data_collection_rule#event_hub_direct MonitorDataCollectionRule#event_hub_direct}
	EventHubDirect *MonitorDataCollectionRuleDestinationsEventHubDirect `field:"optional" json:"eventHubDirect" yaml:"eventHubDirect"`
	// log_analytics block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.49.0/docs/resources/monitor_data_collection_rule#log_analytics MonitorDataCollectionRule#log_analytics}
	LogAnalytics interface{} `field:"optional" json:"logAnalytics" yaml:"logAnalytics"`
	// monitor_account block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.49.0/docs/resources/monitor_data_collection_rule#monitor_account MonitorDataCollectionRule#monitor_account}
	MonitorAccount interface{} `field:"optional" json:"monitorAccount" yaml:"monitorAccount"`
	// storage_blob block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.49.0/docs/resources/monitor_data_collection_rule#storage_blob MonitorDataCollectionRule#storage_blob}
	StorageBlob interface{} `field:"optional" json:"storageBlob" yaml:"storageBlob"`
	// storage_blob_direct block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.49.0/docs/resources/monitor_data_collection_rule#storage_blob_direct MonitorDataCollectionRule#storage_blob_direct}
	StorageBlobDirect interface{} `field:"optional" json:"storageBlobDirect" yaml:"storageBlobDirect"`
	// storage_table_direct block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.49.0/docs/resources/monitor_data_collection_rule#storage_table_direct MonitorDataCollectionRule#storage_table_direct}
	StorageTableDirect interface{} `field:"optional" json:"storageTableDirect" yaml:"storageTableDirect"`
}

