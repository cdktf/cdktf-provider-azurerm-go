// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package storagedatalakegen2filesystem


type StorageDataLakeGen2FilesystemTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.80.0/docs/resources/storage_data_lake_gen2_filesystem#create StorageDataLakeGen2Filesystem#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.80.0/docs/resources/storage_data_lake_gen2_filesystem#delete StorageDataLakeGen2Filesystem#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.80.0/docs/resources/storage_data_lake_gen2_filesystem#read StorageDataLakeGen2Filesystem#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.80.0/docs/resources/storage_data_lake_gen2_filesystem#update StorageDataLakeGen2Filesystem#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

