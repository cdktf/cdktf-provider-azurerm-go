// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package monitordatacollectionrule


type MonitorDataCollectionRuleStreamDeclaration struct {
	// column block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/monitor_data_collection_rule#column MonitorDataCollectionRule#column}
	Column interface{} `field:"required" json:"column" yaml:"column"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/monitor_data_collection_rule#stream_name MonitorDataCollectionRule#stream_name}.
	StreamName *string `field:"required" json:"streamName" yaml:"streamName"`
}

