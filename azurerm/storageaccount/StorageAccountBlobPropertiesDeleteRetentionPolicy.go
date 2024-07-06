// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package storageaccount


type StorageAccountBlobPropertiesDeleteRetentionPolicy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.111.0/docs/resources/storage_account#days StorageAccount#days}.
	Days *float64 `field:"optional" json:"days" yaml:"days"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.111.0/docs/resources/storage_account#permanent_delete_enabled StorageAccount#permanent_delete_enabled}.
	PermanentDeleteEnabled interface{} `field:"optional" json:"permanentDeleteEnabled" yaml:"permanentDeleteEnabled"`
}

