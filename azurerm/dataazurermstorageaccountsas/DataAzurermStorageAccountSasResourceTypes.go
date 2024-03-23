// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataazurermstorageaccountsas


type DataAzurermStorageAccountSasResourceTypes struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.97.1/docs/data-sources/storage_account_sas#container DataAzurermStorageAccountSas#container}.
	Container interface{} `field:"required" json:"container" yaml:"container"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.97.1/docs/data-sources/storage_account_sas#object DataAzurermStorageAccountSas#object}.
	Object interface{} `field:"required" json:"object" yaml:"object"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.97.1/docs/data-sources/storage_account_sas#service DataAzurermStorageAccountSas#service}.
	Service interface{} `field:"required" json:"service" yaml:"service"`
}

