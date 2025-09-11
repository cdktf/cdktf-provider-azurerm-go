// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider


type AzurermProviderFeaturesKeyVault struct {
	// When enabled soft-deleted `azurerm_key_vault_certificate` resources will be permanently deleted (e.g purged), when destroyed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs#purge_soft_deleted_certificates_on_destroy AzurermProvider#purge_soft_deleted_certificates_on_destroy}
	PurgeSoftDeletedCertificatesOnDestroy interface{} `field:"optional" json:"purgeSoftDeletedCertificatesOnDestroy" yaml:"purgeSoftDeletedCertificatesOnDestroy"`
	// When enabled soft-deleted `azurerm_key_vault_managed_hardware_security_module_key` resources will be permanently deleted (e.g purged), when destroyed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs#purge_soft_deleted_hardware_security_module_keys_on_destroy AzurermProvider#purge_soft_deleted_hardware_security_module_keys_on_destroy}
	PurgeSoftDeletedHardwareSecurityModuleKeysOnDestroy interface{} `field:"optional" json:"purgeSoftDeletedHardwareSecurityModuleKeysOnDestroy" yaml:"purgeSoftDeletedHardwareSecurityModuleKeysOnDestroy"`
	// When enabled soft-deleted `azurerm_key_vault_managed_hardware_security_module` resources will be permanently deleted (e.g purged), when destroyed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs#purge_soft_deleted_hardware_security_modules_on_destroy AzurermProvider#purge_soft_deleted_hardware_security_modules_on_destroy}
	PurgeSoftDeletedHardwareSecurityModulesOnDestroy interface{} `field:"optional" json:"purgeSoftDeletedHardwareSecurityModulesOnDestroy" yaml:"purgeSoftDeletedHardwareSecurityModulesOnDestroy"`
	// When enabled soft-deleted `azurerm_key_vault_key` resources will be permanently deleted (e.g purged), when destroyed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs#purge_soft_deleted_keys_on_destroy AzurermProvider#purge_soft_deleted_keys_on_destroy}
	PurgeSoftDeletedKeysOnDestroy interface{} `field:"optional" json:"purgeSoftDeletedKeysOnDestroy" yaml:"purgeSoftDeletedKeysOnDestroy"`
	// When enabled soft-deleted `azurerm_key_vault_secret` resources will be permanently deleted (e.g purged), when destroyed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs#purge_soft_deleted_secrets_on_destroy AzurermProvider#purge_soft_deleted_secrets_on_destroy}
	PurgeSoftDeletedSecretsOnDestroy interface{} `field:"optional" json:"purgeSoftDeletedSecretsOnDestroy" yaml:"purgeSoftDeletedSecretsOnDestroy"`
	// When enabled soft-deleted `azurerm_key_vault` resources will be permanently deleted (e.g purged), when destroyed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs#purge_soft_delete_on_destroy AzurermProvider#purge_soft_delete_on_destroy}
	PurgeSoftDeleteOnDestroy interface{} `field:"optional" json:"purgeSoftDeleteOnDestroy" yaml:"purgeSoftDeleteOnDestroy"`
	// When enabled soft-deleted `azurerm_key_vault_certificate` resources will be restored, instead of creating new ones.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs#recover_soft_deleted_certificates AzurermProvider#recover_soft_deleted_certificates}
	RecoverSoftDeletedCertificates interface{} `field:"optional" json:"recoverSoftDeletedCertificates" yaml:"recoverSoftDeletedCertificates"`
	// When enabled soft-deleted `azurerm_key_vault_managed_hardware_security_module_key` resources will be restored, instead of creating new ones.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs#recover_soft_deleted_hardware_security_module_keys AzurermProvider#recover_soft_deleted_hardware_security_module_keys}
	RecoverSoftDeletedHardwareSecurityModuleKeys interface{} `field:"optional" json:"recoverSoftDeletedHardwareSecurityModuleKeys" yaml:"recoverSoftDeletedHardwareSecurityModuleKeys"`
	// When enabled soft-deleted `azurerm_key_vault_key` resources will be restored, instead of creating new ones.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs#recover_soft_deleted_keys AzurermProvider#recover_soft_deleted_keys}
	RecoverSoftDeletedKeys interface{} `field:"optional" json:"recoverSoftDeletedKeys" yaml:"recoverSoftDeletedKeys"`
	// When enabled soft-deleted `azurerm_key_vault` resources will be restored, instead of creating new ones.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs#recover_soft_deleted_key_vaults AzurermProvider#recover_soft_deleted_key_vaults}
	RecoverSoftDeletedKeyVaults interface{} `field:"optional" json:"recoverSoftDeletedKeyVaults" yaml:"recoverSoftDeletedKeyVaults"`
	// When enabled soft-deleted `azurerm_key_vault_secret` resources will be restored, instead of creating new ones.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs#recover_soft_deleted_secrets AzurermProvider#recover_soft_deleted_secrets}
	RecoverSoftDeletedSecrets interface{} `field:"optional" json:"recoverSoftDeletedSecrets" yaml:"recoverSoftDeletedSecrets"`
}

