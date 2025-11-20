// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package keyvaultmanagedstorageaccountsastokendefinition


type KeyVaultManagedStorageAccountSasTokenDefinitionTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/key_vault_managed_storage_account_sas_token_definition#create KeyVaultManagedStorageAccountSasTokenDefinition#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/key_vault_managed_storage_account_sas_token_definition#delete KeyVaultManagedStorageAccountSasTokenDefinition#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/key_vault_managed_storage_account_sas_token_definition#read KeyVaultManagedStorageAccountSasTokenDefinition#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/key_vault_managed_storage_account_sas_token_definition#update KeyVaultManagedStorageAccountSasTokenDefinition#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

