// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datafactorylinkedserviceazurefilestorage


type DataFactoryLinkedServiceAzureFileStorageKeyVaultPassword struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.28.0/docs/resources/data_factory_linked_service_azure_file_storage#linked_service_name DataFactoryLinkedServiceAzureFileStorage#linked_service_name}.
	LinkedServiceName *string `field:"required" json:"linkedServiceName" yaml:"linkedServiceName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.28.0/docs/resources/data_factory_linked_service_azure_file_storage#secret_name DataFactoryLinkedServiceAzureFileStorage#secret_name}.
	SecretName *string `field:"required" json:"secretName" yaml:"secretName"`
}

