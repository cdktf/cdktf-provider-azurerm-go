// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datafactorylinkedservicepostgresql


type DataFactoryLinkedServicePostgresqlTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.81.0/docs/resources/data_factory_linked_service_postgresql#create DataFactoryLinkedServicePostgresql#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.81.0/docs/resources/data_factory_linked_service_postgresql#delete DataFactoryLinkedServicePostgresql#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.81.0/docs/resources/data_factory_linked_service_postgresql#read DataFactoryLinkedServicePostgresql#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.81.0/docs/resources/data_factory_linked_service_postgresql#update DataFactoryLinkedServicePostgresql#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

