// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package monitordatacollectionrule


type MonitorDataCollectionRuleDataSourcesDataImport struct {
	// event_hub_data_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.87.0/docs/resources/monitor_data_collection_rule#event_hub_data_source MonitorDataCollectionRule#event_hub_data_source}
	EventHubDataSource interface{} `field:"required" json:"eventHubDataSource" yaml:"eventHubDataSource"`
}

