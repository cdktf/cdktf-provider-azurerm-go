// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datafactorylinkedserviceazurefunction


type DataFactoryLinkedServiceAzureFunctionTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.41.0/docs/resources/data_factory_linked_service_azure_function#create DataFactoryLinkedServiceAzureFunction#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.41.0/docs/resources/data_factory_linked_service_azure_function#delete DataFactoryLinkedServiceAzureFunction#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.41.0/docs/resources/data_factory_linked_service_azure_function#read DataFactoryLinkedServiceAzureFunction#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.41.0/docs/resources/data_factory_linked_service_azure_function#update DataFactoryLinkedServiceAzureFunction#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

