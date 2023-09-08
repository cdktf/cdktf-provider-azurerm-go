// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datafactorydatasetbinary

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataFactoryDatasetBinaryConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.72.0/docs/resources/data_factory_dataset_binary#data_factory_id DataFactoryDatasetBinary#data_factory_id}.
	DataFactoryId *string `field:"required" json:"dataFactoryId" yaml:"dataFactoryId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.72.0/docs/resources/data_factory_dataset_binary#linked_service_name DataFactoryDatasetBinary#linked_service_name}.
	LinkedServiceName *string `field:"required" json:"linkedServiceName" yaml:"linkedServiceName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.72.0/docs/resources/data_factory_dataset_binary#name DataFactoryDatasetBinary#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.72.0/docs/resources/data_factory_dataset_binary#additional_properties DataFactoryDatasetBinary#additional_properties}.
	AdditionalProperties *map[string]*string `field:"optional" json:"additionalProperties" yaml:"additionalProperties"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.72.0/docs/resources/data_factory_dataset_binary#annotations DataFactoryDatasetBinary#annotations}.
	Annotations *[]*string `field:"optional" json:"annotations" yaml:"annotations"`
	// azure_blob_storage_location block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.72.0/docs/resources/data_factory_dataset_binary#azure_blob_storage_location DataFactoryDatasetBinary#azure_blob_storage_location}
	AzureBlobStorageLocation *DataFactoryDatasetBinaryAzureBlobStorageLocation `field:"optional" json:"azureBlobStorageLocation" yaml:"azureBlobStorageLocation"`
	// compression block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.72.0/docs/resources/data_factory_dataset_binary#compression DataFactoryDatasetBinary#compression}
	Compression *DataFactoryDatasetBinaryCompression `field:"optional" json:"compression" yaml:"compression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.72.0/docs/resources/data_factory_dataset_binary#description DataFactoryDatasetBinary#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.72.0/docs/resources/data_factory_dataset_binary#folder DataFactoryDatasetBinary#folder}.
	Folder *string `field:"optional" json:"folder" yaml:"folder"`
	// http_server_location block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.72.0/docs/resources/data_factory_dataset_binary#http_server_location DataFactoryDatasetBinary#http_server_location}
	HttpServerLocation *DataFactoryDatasetBinaryHttpServerLocation `field:"optional" json:"httpServerLocation" yaml:"httpServerLocation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.72.0/docs/resources/data_factory_dataset_binary#id DataFactoryDatasetBinary#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.72.0/docs/resources/data_factory_dataset_binary#parameters DataFactoryDatasetBinary#parameters}.
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
	// sftp_server_location block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.72.0/docs/resources/data_factory_dataset_binary#sftp_server_location DataFactoryDatasetBinary#sftp_server_location}
	SftpServerLocation *DataFactoryDatasetBinarySftpServerLocation `field:"optional" json:"sftpServerLocation" yaml:"sftpServerLocation"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.72.0/docs/resources/data_factory_dataset_binary#timeouts DataFactoryDatasetBinary#timeouts}
	Timeouts *DataFactoryDatasetBinaryTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

