// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package iothubendpointcosmosdbaccount

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IothubEndpointCosmosdbAccountConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.82.0/docs/resources/iothub_endpoint_cosmosdb_account#container_name IothubEndpointCosmosdbAccount#container_name}.
	ContainerName *string `field:"required" json:"containerName" yaml:"containerName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.82.0/docs/resources/iothub_endpoint_cosmosdb_account#database_name IothubEndpointCosmosdbAccount#database_name}.
	DatabaseName *string `field:"required" json:"databaseName" yaml:"databaseName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.82.0/docs/resources/iothub_endpoint_cosmosdb_account#endpoint_uri IothubEndpointCosmosdbAccount#endpoint_uri}.
	EndpointUri *string `field:"required" json:"endpointUri" yaml:"endpointUri"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.82.0/docs/resources/iothub_endpoint_cosmosdb_account#iothub_id IothubEndpointCosmosdbAccount#iothub_id}.
	IothubId *string `field:"required" json:"iothubId" yaml:"iothubId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.82.0/docs/resources/iothub_endpoint_cosmosdb_account#name IothubEndpointCosmosdbAccount#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.82.0/docs/resources/iothub_endpoint_cosmosdb_account#resource_group_name IothubEndpointCosmosdbAccount#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.82.0/docs/resources/iothub_endpoint_cosmosdb_account#authentication_type IothubEndpointCosmosdbAccount#authentication_type}.
	AuthenticationType *string `field:"optional" json:"authenticationType" yaml:"authenticationType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.82.0/docs/resources/iothub_endpoint_cosmosdb_account#id IothubEndpointCosmosdbAccount#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.82.0/docs/resources/iothub_endpoint_cosmosdb_account#identity_id IothubEndpointCosmosdbAccount#identity_id}.
	IdentityId *string `field:"optional" json:"identityId" yaml:"identityId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.82.0/docs/resources/iothub_endpoint_cosmosdb_account#partition_key_name IothubEndpointCosmosdbAccount#partition_key_name}.
	PartitionKeyName *string `field:"optional" json:"partitionKeyName" yaml:"partitionKeyName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.82.0/docs/resources/iothub_endpoint_cosmosdb_account#partition_key_template IothubEndpointCosmosdbAccount#partition_key_template}.
	PartitionKeyTemplate *string `field:"optional" json:"partitionKeyTemplate" yaml:"partitionKeyTemplate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.82.0/docs/resources/iothub_endpoint_cosmosdb_account#primary_key IothubEndpointCosmosdbAccount#primary_key}.
	PrimaryKey *string `field:"optional" json:"primaryKey" yaml:"primaryKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.82.0/docs/resources/iothub_endpoint_cosmosdb_account#secondary_key IothubEndpointCosmosdbAccount#secondary_key}.
	SecondaryKey *string `field:"optional" json:"secondaryKey" yaml:"secondaryKey"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.82.0/docs/resources/iothub_endpoint_cosmosdb_account#timeouts IothubEndpointCosmosdbAccount#timeouts}
	Timeouts *IothubEndpointCosmosdbAccountTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

