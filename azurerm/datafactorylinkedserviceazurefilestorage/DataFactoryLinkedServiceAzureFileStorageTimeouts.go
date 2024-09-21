// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datafactorylinkedserviceazurefilestorage


type DataFactoryLinkedServiceAzureFileStorageTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.3.0/docs/resources/data_factory_linked_service_azure_file_storage#create DataFactoryLinkedServiceAzureFileStorage#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.3.0/docs/resources/data_factory_linked_service_azure_file_storage#delete DataFactoryLinkedServiceAzureFileStorage#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.3.0/docs/resources/data_factory_linked_service_azure_file_storage#read DataFactoryLinkedServiceAzureFileStorage#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.3.0/docs/resources/data_factory_linked_service_azure_file_storage#update DataFactoryLinkedServiceAzureFileStorage#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

