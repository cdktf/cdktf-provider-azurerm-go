// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datafactorylinkedservicemysql


type DataFactoryLinkedServiceMysqlTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/data_factory_linked_service_mysql#create DataFactoryLinkedServiceMysql#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/data_factory_linked_service_mysql#delete DataFactoryLinkedServiceMysql#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/data_factory_linked_service_mysql#read DataFactoryLinkedServiceMysql#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/data_factory_linked_service_mysql#update DataFactoryLinkedServiceMysql#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

