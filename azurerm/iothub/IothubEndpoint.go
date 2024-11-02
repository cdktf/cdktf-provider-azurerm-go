// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package iothub


type IothubEndpoint struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/iothub#authentication_type Iothub#authentication_type}.
	AuthenticationType *string `field:"optional" json:"authenticationType" yaml:"authenticationType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/iothub#batch_frequency_in_seconds Iothub#batch_frequency_in_seconds}.
	BatchFrequencyInSeconds *float64 `field:"optional" json:"batchFrequencyInSeconds" yaml:"batchFrequencyInSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/iothub#connection_string Iothub#connection_string}.
	ConnectionString *string `field:"optional" json:"connectionString" yaml:"connectionString"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/iothub#container_name Iothub#container_name}.
	ContainerName *string `field:"optional" json:"containerName" yaml:"containerName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/iothub#encoding Iothub#encoding}.
	Encoding *string `field:"optional" json:"encoding" yaml:"encoding"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/iothub#endpoint_uri Iothub#endpoint_uri}.
	EndpointUri *string `field:"optional" json:"endpointUri" yaml:"endpointUri"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/iothub#entity_path Iothub#entity_path}.
	EntityPath *string `field:"optional" json:"entityPath" yaml:"entityPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/iothub#file_name_format Iothub#file_name_format}.
	FileNameFormat *string `field:"optional" json:"fileNameFormat" yaml:"fileNameFormat"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/iothub#identity_id Iothub#identity_id}.
	IdentityId *string `field:"optional" json:"identityId" yaml:"identityId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/iothub#max_chunk_size_in_bytes Iothub#max_chunk_size_in_bytes}.
	MaxChunkSizeInBytes *float64 `field:"optional" json:"maxChunkSizeInBytes" yaml:"maxChunkSizeInBytes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/iothub#name Iothub#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/iothub#resource_group_name Iothub#resource_group_name}.
	ResourceGroupName *string `field:"optional" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/iothub#type Iothub#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

