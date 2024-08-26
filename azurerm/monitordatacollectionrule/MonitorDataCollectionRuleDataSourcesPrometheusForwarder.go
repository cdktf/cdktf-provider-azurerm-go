// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package monitordatacollectionrule


type MonitorDataCollectionRuleDataSourcesPrometheusForwarder struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.0.1/docs/resources/monitor_data_collection_rule#name MonitorDataCollectionRule#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.0.1/docs/resources/monitor_data_collection_rule#streams MonitorDataCollectionRule#streams}.
	Streams *[]*string `field:"required" json:"streams" yaml:"streams"`
	// label_include_filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.0.1/docs/resources/monitor_data_collection_rule#label_include_filter MonitorDataCollectionRule#label_include_filter}
	LabelIncludeFilter interface{} `field:"optional" json:"labelIncludeFilter" yaml:"labelIncludeFilter"`
}

