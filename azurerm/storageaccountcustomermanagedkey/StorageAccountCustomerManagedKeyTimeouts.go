// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package storageaccountcustomermanagedkey


type StorageAccountCustomerManagedKeyTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.0/docs/resources/storage_account_customer_managed_key#create StorageAccountCustomerManagedKeyA#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.0/docs/resources/storage_account_customer_managed_key#delete StorageAccountCustomerManagedKeyA#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.0/docs/resources/storage_account_customer_managed_key#read StorageAccountCustomerManagedKeyA#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.0/docs/resources/storage_account_customer_managed_key#update StorageAccountCustomerManagedKeyA#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

