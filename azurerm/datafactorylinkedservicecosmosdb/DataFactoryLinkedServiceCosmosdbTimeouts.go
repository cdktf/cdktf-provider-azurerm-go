// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datafactorylinkedservicecosmosdb


type DataFactoryLinkedServiceCosmosdbTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/data_factory_linked_service_cosmosdb#create DataFactoryLinkedServiceCosmosdb#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/data_factory_linked_service_cosmosdb#delete DataFactoryLinkedServiceCosmosdb#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/data_factory_linked_service_cosmosdb#read DataFactoryLinkedServiceCosmosdb#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/data_factory_linked_service_cosmosdb#update DataFactoryLinkedServiceCosmosdb#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

