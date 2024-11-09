// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package storageaccount


type StorageAccountAzureFilesAuthenticationActiveDirectory struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.9.0/docs/resources/storage_account#domain_guid StorageAccount#domain_guid}.
	DomainGuid *string `field:"required" json:"domainGuid" yaml:"domainGuid"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.9.0/docs/resources/storage_account#domain_name StorageAccount#domain_name}.
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.9.0/docs/resources/storage_account#domain_sid StorageAccount#domain_sid}.
	DomainSid *string `field:"optional" json:"domainSid" yaml:"domainSid"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.9.0/docs/resources/storage_account#forest_name StorageAccount#forest_name}.
	ForestName *string `field:"optional" json:"forestName" yaml:"forestName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.9.0/docs/resources/storage_account#netbios_domain_name StorageAccount#netbios_domain_name}.
	NetbiosDomainName *string `field:"optional" json:"netbiosDomainName" yaml:"netbiosDomainName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.9.0/docs/resources/storage_account#storage_sid StorageAccount#storage_sid}.
	StorageSid *string `field:"optional" json:"storageSid" yaml:"storageSid"`
}

