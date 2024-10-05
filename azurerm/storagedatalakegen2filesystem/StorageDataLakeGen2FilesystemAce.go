// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package storagedatalakegen2filesystem


type StorageDataLakeGen2FilesystemAce struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.4.0/docs/resources/storage_data_lake_gen2_filesystem#permissions StorageDataLakeGen2Filesystem#permissions}.
	Permissions *string `field:"required" json:"permissions" yaml:"permissions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.4.0/docs/resources/storage_data_lake_gen2_filesystem#type StorageDataLakeGen2Filesystem#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.4.0/docs/resources/storage_data_lake_gen2_filesystem#id StorageDataLakeGen2Filesystem#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.4.0/docs/resources/storage_data_lake_gen2_filesystem#scope StorageDataLakeGen2Filesystem#scope}.
	Scope *string `field:"optional" json:"scope" yaml:"scope"`
}

