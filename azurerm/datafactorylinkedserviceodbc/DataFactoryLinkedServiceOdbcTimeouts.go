// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datafactorylinkedserviceodbc


type DataFactoryLinkedServiceOdbcTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.100.0/docs/resources/data_factory_linked_service_odbc#create DataFactoryLinkedServiceOdbc#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.100.0/docs/resources/data_factory_linked_service_odbc#delete DataFactoryLinkedServiceOdbc#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.100.0/docs/resources/data_factory_linked_service_odbc#read DataFactoryLinkedServiceOdbc#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.100.0/docs/resources/data_factory_linked_service_odbc#update DataFactoryLinkedServiceOdbc#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

