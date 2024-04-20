// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package storageaccount


type StorageAccountSasPolicy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.100.0/docs/resources/storage_account#expiration_period StorageAccount#expiration_period}.
	ExpirationPeriod *string `field:"required" json:"expirationPeriod" yaml:"expirationPeriod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.100.0/docs/resources/storage_account#expiration_action StorageAccount#expiration_action}.
	ExpirationAction *string `field:"optional" json:"expirationAction" yaml:"expirationAction"`
}

