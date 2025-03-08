// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package monitordatacollectionrule


type MonitorDataCollectionRuleDataSourcesExtension struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.22.0/docs/resources/monitor_data_collection_rule#extension_name MonitorDataCollectionRule#extension_name}.
	ExtensionName *string `field:"required" json:"extensionName" yaml:"extensionName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.22.0/docs/resources/monitor_data_collection_rule#name MonitorDataCollectionRule#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.22.0/docs/resources/monitor_data_collection_rule#streams MonitorDataCollectionRule#streams}.
	Streams *[]*string `field:"required" json:"streams" yaml:"streams"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.22.0/docs/resources/monitor_data_collection_rule#extension_json MonitorDataCollectionRule#extension_json}.
	ExtensionJson *string `field:"optional" json:"extensionJson" yaml:"extensionJson"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.22.0/docs/resources/monitor_data_collection_rule#input_data_sources MonitorDataCollectionRule#input_data_sources}.
	InputDataSources *[]*string `field:"optional" json:"inputDataSources" yaml:"inputDataSources"`
}

