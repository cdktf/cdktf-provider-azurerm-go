// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containerregistrycredentialset


type ContainerRegistryCredentialSetAuthenticationCredentials struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.40.0/docs/resources/container_registry_credential_set#password_secret_id ContainerRegistryCredentialSet#password_secret_id}.
	PasswordSecretId *string `field:"required" json:"passwordSecretId" yaml:"passwordSecretId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.40.0/docs/resources/container_registry_credential_set#username_secret_id ContainerRegistryCredentialSet#username_secret_id}.
	UsernameSecretId *string `field:"required" json:"usernameSecretId" yaml:"usernameSecretId"`
}

