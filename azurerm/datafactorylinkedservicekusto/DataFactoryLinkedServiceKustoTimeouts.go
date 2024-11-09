// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datafactorylinkedservicekusto


type DataFactoryLinkedServiceKustoTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.9.0/docs/resources/data_factory_linked_service_kusto#create DataFactoryLinkedServiceKusto#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.9.0/docs/resources/data_factory_linked_service_kusto#delete DataFactoryLinkedServiceKusto#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.9.0/docs/resources/data_factory_linked_service_kusto#read DataFactoryLinkedServiceKusto#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.9.0/docs/resources/data_factory_linked_service_kusto#update DataFactoryLinkedServiceKusto#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

