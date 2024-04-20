// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datafactorycredentialusermanagedidentity


type DataFactoryCredentialUserManagedIdentityTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.100.0/docs/resources/data_factory_credential_user_managed_identity#create DataFactoryCredentialUserManagedIdentity#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.100.0/docs/resources/data_factory_credential_user_managed_identity#delete DataFactoryCredentialUserManagedIdentity#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.100.0/docs/resources/data_factory_credential_user_managed_identity#read DataFactoryCredentialUserManagedIdentity#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.100.0/docs/resources/data_factory_credential_user_managed_identity#update DataFactoryCredentialUserManagedIdentity#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

