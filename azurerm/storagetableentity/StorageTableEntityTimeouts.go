// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package storagetableentity


type StorageTableEntityTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/storage_table_entity#create StorageTableEntity#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/storage_table_entity#delete StorageTableEntity#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/storage_table_entity#read StorageTableEntity#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/storage_table_entity#update StorageTableEntity#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

