// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package keyvault


type KeyVaultAccessPolicy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/key_vault#application_id KeyVault#application_id}.
	ApplicationId *string `field:"optional" json:"applicationId" yaml:"applicationId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/key_vault#certificate_permissions KeyVault#certificate_permissions}.
	CertificatePermissions *[]*string `field:"optional" json:"certificatePermissions" yaml:"certificatePermissions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/key_vault#key_permissions KeyVault#key_permissions}.
	KeyPermissions *[]*string `field:"optional" json:"keyPermissions" yaml:"keyPermissions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/key_vault#object_id KeyVault#object_id}.
	ObjectId *string `field:"optional" json:"objectId" yaml:"objectId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/key_vault#secret_permissions KeyVault#secret_permissions}.
	SecretPermissions *[]*string `field:"optional" json:"secretPermissions" yaml:"secretPermissions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/key_vault#storage_permissions KeyVault#storage_permissions}.
	StoragePermissions *[]*string `field:"optional" json:"storagePermissions" yaml:"storagePermissions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/key_vault#tenant_id KeyVault#tenant_id}.
	TenantId *string `field:"optional" json:"tenantId" yaml:"tenantId"`
}

