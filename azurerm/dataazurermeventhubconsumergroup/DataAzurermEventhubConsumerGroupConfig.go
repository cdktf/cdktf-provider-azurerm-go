// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataazurermeventhubconsumergroup

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataAzurermEventhubConsumerGroupConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/data-sources/eventhub_consumer_group#eventhub_name DataAzurermEventhubConsumerGroup#eventhub_name}.
	EventhubName *string `field:"required" json:"eventhubName" yaml:"eventhubName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/data-sources/eventhub_consumer_group#name DataAzurermEventhubConsumerGroup#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/data-sources/eventhub_consumer_group#namespace_name DataAzurermEventhubConsumerGroup#namespace_name}.
	NamespaceName *string `field:"required" json:"namespaceName" yaml:"namespaceName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/data-sources/eventhub_consumer_group#resource_group_name DataAzurermEventhubConsumerGroup#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/data-sources/eventhub_consumer_group#id DataAzurermEventhubConsumerGroup#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/data-sources/eventhub_consumer_group#timeouts DataAzurermEventhubConsumerGroup#timeouts}
	Timeouts *DataAzurermEventhubConsumerGroupTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

