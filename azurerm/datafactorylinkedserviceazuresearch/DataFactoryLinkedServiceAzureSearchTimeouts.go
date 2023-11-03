// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datafactorylinkedserviceazuresearch


type DataFactoryLinkedServiceAzureSearchTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.79.0/docs/resources/data_factory_linked_service_azure_search#create DataFactoryLinkedServiceAzureSearch#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.79.0/docs/resources/data_factory_linked_service_azure_search#delete DataFactoryLinkedServiceAzureSearch#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.79.0/docs/resources/data_factory_linked_service_azure_search#read DataFactoryLinkedServiceAzureSearch#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.79.0/docs/resources/data_factory_linked_service_azure_search#update DataFactoryLinkedServiceAzureSearch#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

