// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datafactorycredentialserviceprincipal


type DataFactoryCredentialServicePrincipalTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/data_factory_credential_service_principal#create DataFactoryCredentialServicePrincipal#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/data_factory_credential_service_principal#delete DataFactoryCredentialServicePrincipal#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/data_factory_credential_service_principal#read DataFactoryCredentialServicePrincipal#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/data_factory_credential_service_principal#update DataFactoryCredentialServicePrincipal#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

