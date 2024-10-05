// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containerapp


type ContainerAppSecret struct {
	// The secret name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.4.0/docs/resources/container_app#name ContainerApp#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The identity to use for accessing key vault reference.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.4.0/docs/resources/container_app#identity ContainerApp#identity}
	Identity *string `field:"optional" json:"identity" yaml:"identity"`
	// The Key Vault Secret ID. Could be either one of `id` or `versionless_id`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.4.0/docs/resources/container_app#key_vault_secret_id ContainerApp#key_vault_secret_id}
	KeyVaultSecretId *string `field:"optional" json:"keyVaultSecretId" yaml:"keyVaultSecretId"`
	// The value for this secret.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.4.0/docs/resources/container_app#value ContainerApp#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

