// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package monitordatacollectionruleassociation


type MonitorDataCollectionRuleAssociationTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.41.0/docs/resources/monitor_data_collection_rule_association#create MonitorDataCollectionRuleAssociation#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.41.0/docs/resources/monitor_data_collection_rule_association#delete MonitorDataCollectionRuleAssociation#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.41.0/docs/resources/monitor_data_collection_rule_association#read MonitorDataCollectionRuleAssociation#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.41.0/docs/resources/monitor_data_collection_rule_association#update MonitorDataCollectionRuleAssociation#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

