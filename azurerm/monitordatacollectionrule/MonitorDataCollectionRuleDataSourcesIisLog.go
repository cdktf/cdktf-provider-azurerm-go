// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package monitordatacollectionrule


type MonitorDataCollectionRuleDataSourcesIisLog struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.102.0/docs/resources/monitor_data_collection_rule#name MonitorDataCollectionRule#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.102.0/docs/resources/monitor_data_collection_rule#streams MonitorDataCollectionRule#streams}.
	Streams *[]*string `field:"required" json:"streams" yaml:"streams"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.102.0/docs/resources/monitor_data_collection_rule#log_directories MonitorDataCollectionRule#log_directories}.
	LogDirectories *[]*string `field:"optional" json:"logDirectories" yaml:"logDirectories"`
}

