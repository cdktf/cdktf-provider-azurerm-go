// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package iothubendpointstoragecontainer

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IothubEndpointStorageContainerConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/iothub_endpoint_storage_container#container_name IothubEndpointStorageContainer#container_name}.
	ContainerName *string `field:"required" json:"containerName" yaml:"containerName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/iothub_endpoint_storage_container#iothub_id IothubEndpointStorageContainer#iothub_id}.
	IothubId *string `field:"required" json:"iothubId" yaml:"iothubId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/iothub_endpoint_storage_container#name IothubEndpointStorageContainer#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/iothub_endpoint_storage_container#resource_group_name IothubEndpointStorageContainer#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/iothub_endpoint_storage_container#authentication_type IothubEndpointStorageContainer#authentication_type}.
	AuthenticationType *string `field:"optional" json:"authenticationType" yaml:"authenticationType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/iothub_endpoint_storage_container#batch_frequency_in_seconds IothubEndpointStorageContainer#batch_frequency_in_seconds}.
	BatchFrequencyInSeconds *float64 `field:"optional" json:"batchFrequencyInSeconds" yaml:"batchFrequencyInSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/iothub_endpoint_storage_container#connection_string IothubEndpointStorageContainer#connection_string}.
	ConnectionString *string `field:"optional" json:"connectionString" yaml:"connectionString"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/iothub_endpoint_storage_container#encoding IothubEndpointStorageContainer#encoding}.
	Encoding *string `field:"optional" json:"encoding" yaml:"encoding"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/iothub_endpoint_storage_container#endpoint_uri IothubEndpointStorageContainer#endpoint_uri}.
	EndpointUri *string `field:"optional" json:"endpointUri" yaml:"endpointUri"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/iothub_endpoint_storage_container#file_name_format IothubEndpointStorageContainer#file_name_format}.
	FileNameFormat *string `field:"optional" json:"fileNameFormat" yaml:"fileNameFormat"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/iothub_endpoint_storage_container#id IothubEndpointStorageContainer#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/iothub_endpoint_storage_container#identity_id IothubEndpointStorageContainer#identity_id}.
	IdentityId *string `field:"optional" json:"identityId" yaml:"identityId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/iothub_endpoint_storage_container#max_chunk_size_in_bytes IothubEndpointStorageContainer#max_chunk_size_in_bytes}.
	MaxChunkSizeInBytes *float64 `field:"optional" json:"maxChunkSizeInBytes" yaml:"maxChunkSizeInBytes"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/iothub_endpoint_storage_container#timeouts IothubEndpointStorageContainer#timeouts}
	Timeouts *IothubEndpointStorageContainerTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

