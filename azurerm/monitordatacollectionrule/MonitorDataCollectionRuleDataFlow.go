// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package monitordatacollectionrule


type MonitorDataCollectionRuleDataFlow struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.111.0/docs/resources/monitor_data_collection_rule#destinations MonitorDataCollectionRule#destinations}.
	Destinations *[]*string `field:"required" json:"destinations" yaml:"destinations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.111.0/docs/resources/monitor_data_collection_rule#streams MonitorDataCollectionRule#streams}.
	Streams *[]*string `field:"required" json:"streams" yaml:"streams"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.111.0/docs/resources/monitor_data_collection_rule#built_in_transform MonitorDataCollectionRule#built_in_transform}.
	BuiltInTransform *string `field:"optional" json:"builtInTransform" yaml:"builtInTransform"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.111.0/docs/resources/monitor_data_collection_rule#output_stream MonitorDataCollectionRule#output_stream}.
	OutputStream *string `field:"optional" json:"outputStream" yaml:"outputStream"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.111.0/docs/resources/monitor_data_collection_rule#transform_kql MonitorDataCollectionRule#transform_kql}.
	TransformKql *string `field:"optional" json:"transformKql" yaml:"transformKql"`
}

