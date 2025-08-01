// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package storageaccount


type StorageAccountBlobPropertiesRestorePolicy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/storage_account#days StorageAccount#days}.
	Days *float64 `field:"required" json:"days" yaml:"days"`
}

