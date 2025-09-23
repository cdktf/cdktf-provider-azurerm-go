// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datafactorylinkedservicesftp


type DataFactoryLinkedServiceSftpKeyVaultPassword struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/data_factory_linked_service_sftp#linked_service_name DataFactoryLinkedServiceSftp#linked_service_name}.
	LinkedServiceName *string `field:"required" json:"linkedServiceName" yaml:"linkedServiceName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/data_factory_linked_service_sftp#secret_name DataFactoryLinkedServiceSftp#secret_name}.
	SecretName *string `field:"required" json:"secretName" yaml:"secretName"`
}

