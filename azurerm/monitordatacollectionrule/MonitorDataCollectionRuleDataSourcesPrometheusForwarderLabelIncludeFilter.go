// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package monitordatacollectionrule


type MonitorDataCollectionRuleDataSourcesPrometheusForwarderLabelIncludeFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.4.0/docs/resources/monitor_data_collection_rule#label MonitorDataCollectionRule#label}.
	Label *string `field:"required" json:"label" yaml:"label"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.4.0/docs/resources/monitor_data_collection_rule#value MonitorDataCollectionRule#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

