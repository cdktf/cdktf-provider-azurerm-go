// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package storagedatalakegen2path


type StorageDataLakeGen2PathTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.108.0/docs/resources/storage_data_lake_gen2_path#create StorageDataLakeGen2Path#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.108.0/docs/resources/storage_data_lake_gen2_path#delete StorageDataLakeGen2Path#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.108.0/docs/resources/storage_data_lake_gen2_path#read StorageDataLakeGen2Path#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.108.0/docs/resources/storage_data_lake_gen2_path#update StorageDataLakeGen2Path#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

