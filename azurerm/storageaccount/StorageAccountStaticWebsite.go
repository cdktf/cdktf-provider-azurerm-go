// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package storageaccount


type StorageAccountStaticWebsite struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.50.0/docs/resources/storage_account#error_404_document StorageAccount#error_404_document}.
	Error404Document *string `field:"optional" json:"error404Document" yaml:"error404Document"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.50.0/docs/resources/storage_account#index_document StorageAccount#index_document}.
	IndexDocument *string `field:"optional" json:"indexDocument" yaml:"indexDocument"`
}

