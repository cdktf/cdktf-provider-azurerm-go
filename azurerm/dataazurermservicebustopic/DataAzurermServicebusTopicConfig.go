// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataazurermservicebustopic

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataAzurermServicebusTopicConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/data-sources/servicebus_topic#name DataAzurermServicebusTopic#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/data-sources/servicebus_topic#id DataAzurermServicebusTopic#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/data-sources/servicebus_topic#namespace_id DataAzurermServicebusTopic#namespace_id}.
	NamespaceId *string `field:"optional" json:"namespaceId" yaml:"namespaceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/data-sources/servicebus_topic#namespace_name DataAzurermServicebusTopic#namespace_name}.
	NamespaceName *string `field:"optional" json:"namespaceName" yaml:"namespaceName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/data-sources/servicebus_topic#resource_group_name DataAzurermServicebusTopic#resource_group_name}.
	ResourceGroupName *string `field:"optional" json:"resourceGroupName" yaml:"resourceGroupName"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/data-sources/servicebus_topic#timeouts DataAzurermServicebusTopic#timeouts}
	Timeouts *DataAzurermServicebusTopicTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

