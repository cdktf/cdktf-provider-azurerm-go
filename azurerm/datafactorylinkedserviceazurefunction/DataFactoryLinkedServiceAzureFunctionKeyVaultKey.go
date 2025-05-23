// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datafactorylinkedserviceazurefunction


type DataFactoryLinkedServiceAzureFunctionKeyVaultKey struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.30.0/docs/resources/data_factory_linked_service_azure_function#linked_service_name DataFactoryLinkedServiceAzureFunction#linked_service_name}.
	LinkedServiceName *string `field:"required" json:"linkedServiceName" yaml:"linkedServiceName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.30.0/docs/resources/data_factory_linked_service_azure_function#secret_name DataFactoryLinkedServiceAzureFunction#secret_name}.
	SecretName *string `field:"required" json:"secretName" yaml:"secretName"`
}

