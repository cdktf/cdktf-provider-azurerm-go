// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package monitordatacollectionrule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MonitorDataCollectionRuleConfig struct {
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
	// data_flow block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/resources/monitor_data_collection_rule#data_flow MonitorDataCollectionRule#data_flow}
	DataFlow interface{} `field:"required" json:"dataFlow" yaml:"dataFlow"`
	// destinations block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/resources/monitor_data_collection_rule#destinations MonitorDataCollectionRule#destinations}
	Destinations *MonitorDataCollectionRuleDestinations `field:"required" json:"destinations" yaml:"destinations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/resources/monitor_data_collection_rule#location MonitorDataCollectionRule#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/resources/monitor_data_collection_rule#name MonitorDataCollectionRule#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/resources/monitor_data_collection_rule#resource_group_name MonitorDataCollectionRule#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/resources/monitor_data_collection_rule#data_collection_endpoint_id MonitorDataCollectionRule#data_collection_endpoint_id}.
	DataCollectionEndpointId *string `field:"optional" json:"dataCollectionEndpointId" yaml:"dataCollectionEndpointId"`
	// data_sources block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/resources/monitor_data_collection_rule#data_sources MonitorDataCollectionRule#data_sources}
	DataSources *MonitorDataCollectionRuleDataSources `field:"optional" json:"dataSources" yaml:"dataSources"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/resources/monitor_data_collection_rule#description MonitorDataCollectionRule#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/resources/monitor_data_collection_rule#id MonitorDataCollectionRule#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// identity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/resources/monitor_data_collection_rule#identity MonitorDataCollectionRule#identity}
	Identity *MonitorDataCollectionRuleIdentity `field:"optional" json:"identity" yaml:"identity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/resources/monitor_data_collection_rule#kind MonitorDataCollectionRule#kind}.
	Kind *string `field:"optional" json:"kind" yaml:"kind"`
	// stream_declaration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/resources/monitor_data_collection_rule#stream_declaration MonitorDataCollectionRule#stream_declaration}
	StreamDeclaration interface{} `field:"optional" json:"streamDeclaration" yaml:"streamDeclaration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/resources/monitor_data_collection_rule#tags MonitorDataCollectionRule#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/resources/monitor_data_collection_rule#timeouts MonitorDataCollectionRule#timeouts}
	Timeouts *MonitorDataCollectionRuleTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

