// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package storageaccountqueueproperties


type StorageAccountQueuePropertiesTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/storage_account_queue_properties#create StorageAccountQueuePropertiesA#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/storage_account_queue_properties#delete StorageAccountQueuePropertiesA#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/storage_account_queue_properties#read StorageAccountQueuePropertiesA#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/storage_account_queue_properties#update StorageAccountQueuePropertiesA#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

