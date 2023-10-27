// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datafactorylinkedserviceazuretablestorage


type DataFactoryLinkedServiceAzureTableStorageTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.78.0/docs/resources/data_factory_linked_service_azure_table_storage#create DataFactoryLinkedServiceAzureTableStorage#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.78.0/docs/resources/data_factory_linked_service_azure_table_storage#delete DataFactoryLinkedServiceAzureTableStorage#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.78.0/docs/resources/data_factory_linked_service_azure_table_storage#read DataFactoryLinkedServiceAzureTableStorage#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.78.0/docs/resources/data_factory_linked_service_azure_table_storage#update DataFactoryLinkedServiceAzureTableStorage#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

