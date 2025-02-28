// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package monitordatacollectionruleassociation

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MonitorDataCollectionRuleAssociationConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.0/docs/resources/monitor_data_collection_rule_association#target_resource_id MonitorDataCollectionRuleAssociation#target_resource_id}.
	TargetResourceId *string `field:"required" json:"targetResourceId" yaml:"targetResourceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.0/docs/resources/monitor_data_collection_rule_association#data_collection_endpoint_id MonitorDataCollectionRuleAssociation#data_collection_endpoint_id}.
	DataCollectionEndpointId *string `field:"optional" json:"dataCollectionEndpointId" yaml:"dataCollectionEndpointId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.0/docs/resources/monitor_data_collection_rule_association#data_collection_rule_id MonitorDataCollectionRuleAssociation#data_collection_rule_id}.
	DataCollectionRuleId *string `field:"optional" json:"dataCollectionRuleId" yaml:"dataCollectionRuleId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.0/docs/resources/monitor_data_collection_rule_association#description MonitorDataCollectionRuleAssociation#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.0/docs/resources/monitor_data_collection_rule_association#id MonitorDataCollectionRuleAssociation#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.0/docs/resources/monitor_data_collection_rule_association#name MonitorDataCollectionRuleAssociation#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.0/docs/resources/monitor_data_collection_rule_association#timeouts MonitorDataCollectionRuleAssociation#timeouts}
	Timeouts *MonitorDataCollectionRuleAssociationTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

