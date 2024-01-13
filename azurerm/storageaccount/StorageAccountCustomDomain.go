// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package storageaccount


type StorageAccountCustomDomain struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.87.0/docs/resources/storage_account#name StorageAccount#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.87.0/docs/resources/storage_account#use_subdomain StorageAccount#use_subdomain}.
	UseSubdomain interface{} `field:"optional" json:"useSubdomain" yaml:"useSubdomain"`
}

