// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datasharedatasetblobstorage

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataShareDatasetBlobStorageConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.79.0/docs/resources/data_share_dataset_blob_storage#container_name DataShareDatasetBlobStorage#container_name}.
	ContainerName *string `field:"required" json:"containerName" yaml:"containerName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.79.0/docs/resources/data_share_dataset_blob_storage#data_share_id DataShareDatasetBlobStorage#data_share_id}.
	DataShareId *string `field:"required" json:"dataShareId" yaml:"dataShareId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.79.0/docs/resources/data_share_dataset_blob_storage#name DataShareDatasetBlobStorage#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// storage_account block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.79.0/docs/resources/data_share_dataset_blob_storage#storage_account DataShareDatasetBlobStorage#storage_account}
	StorageAccount *DataShareDatasetBlobStorageStorageAccount `field:"required" json:"storageAccount" yaml:"storageAccount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.79.0/docs/resources/data_share_dataset_blob_storage#file_path DataShareDatasetBlobStorage#file_path}.
	FilePath *string `field:"optional" json:"filePath" yaml:"filePath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.79.0/docs/resources/data_share_dataset_blob_storage#folder_path DataShareDatasetBlobStorage#folder_path}.
	FolderPath *string `field:"optional" json:"folderPath" yaml:"folderPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.79.0/docs/resources/data_share_dataset_blob_storage#id DataShareDatasetBlobStorage#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.79.0/docs/resources/data_share_dataset_blob_storage#timeouts DataShareDatasetBlobStorage#timeouts}
	Timeouts *DataShareDatasetBlobStorageTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

