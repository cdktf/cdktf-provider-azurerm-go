// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datasharedatasetdatalakegen2

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataShareDatasetDataLakeGen2Config struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.111.0/docs/resources/data_share_dataset_data_lake_gen2#file_system_name DataShareDatasetDataLakeGen2#file_system_name}.
	FileSystemName *string `field:"required" json:"fileSystemName" yaml:"fileSystemName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.111.0/docs/resources/data_share_dataset_data_lake_gen2#name DataShareDatasetDataLakeGen2#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.111.0/docs/resources/data_share_dataset_data_lake_gen2#share_id DataShareDatasetDataLakeGen2#share_id}.
	ShareId *string `field:"required" json:"shareId" yaml:"shareId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.111.0/docs/resources/data_share_dataset_data_lake_gen2#storage_account_id DataShareDatasetDataLakeGen2#storage_account_id}.
	StorageAccountId *string `field:"required" json:"storageAccountId" yaml:"storageAccountId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.111.0/docs/resources/data_share_dataset_data_lake_gen2#file_path DataShareDatasetDataLakeGen2#file_path}.
	FilePath *string `field:"optional" json:"filePath" yaml:"filePath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.111.0/docs/resources/data_share_dataset_data_lake_gen2#folder_path DataShareDatasetDataLakeGen2#folder_path}.
	FolderPath *string `field:"optional" json:"folderPath" yaml:"folderPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.111.0/docs/resources/data_share_dataset_data_lake_gen2#id DataShareDatasetDataLakeGen2#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.111.0/docs/resources/data_share_dataset_data_lake_gen2#timeouts DataShareDatasetDataLakeGen2#timeouts}
	Timeouts *DataShareDatasetDataLakeGen2Timeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

