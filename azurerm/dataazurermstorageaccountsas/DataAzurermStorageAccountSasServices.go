// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataazurermstorageaccountsas


type DataAzurermStorageAccountSasServices struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.1/docs/data-sources/storage_account_sas#blob DataAzurermStorageAccountSas#blob}.
	Blob interface{} `field:"required" json:"blob" yaml:"blob"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.1/docs/data-sources/storage_account_sas#file DataAzurermStorageAccountSas#file}.
	File interface{} `field:"required" json:"file" yaml:"file"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.1/docs/data-sources/storage_account_sas#queue DataAzurermStorageAccountSas#queue}.
	Queue interface{} `field:"required" json:"queue" yaml:"queue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.1/docs/data-sources/storage_account_sas#table DataAzurermStorageAccountSas#table}.
	Table interface{} `field:"required" json:"table" yaml:"table"`
}

