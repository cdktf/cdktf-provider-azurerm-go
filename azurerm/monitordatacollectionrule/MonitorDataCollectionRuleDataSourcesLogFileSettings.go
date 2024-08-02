// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package monitordatacollectionrule


type MonitorDataCollectionRuleDataSourcesLogFileSettings struct {
	// text block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.114.0/docs/resources/monitor_data_collection_rule#text MonitorDataCollectionRule#text}
	Text *MonitorDataCollectionRuleDataSourcesLogFileSettingsText `field:"required" json:"text" yaml:"text"`
}

