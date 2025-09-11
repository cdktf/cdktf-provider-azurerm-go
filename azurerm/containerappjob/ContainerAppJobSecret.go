// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containerappjob


type ContainerAppJobSecret struct {
	// The secret name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs/resources/container_app_job#name ContainerAppJob#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The identity to use for accessing key vault reference.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs/resources/container_app_job#identity ContainerAppJob#identity}
	Identity *string `field:"optional" json:"identity" yaml:"identity"`
	// The Key Vault Secret ID. Could be either one of `id` or `versionless_id`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs/resources/container_app_job#key_vault_secret_id ContainerAppJob#key_vault_secret_id}
	KeyVaultSecretId *string `field:"optional" json:"keyVaultSecretId" yaml:"keyVaultSecretId"`
	// The value for this secret.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs/resources/container_app_job#value ContainerAppJob#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

