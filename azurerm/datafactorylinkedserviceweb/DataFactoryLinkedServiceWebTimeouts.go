// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datafactorylinkedserviceweb


type DataFactoryLinkedServiceWebTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.24.0/docs/resources/data_factory_linked_service_web#create DataFactoryLinkedServiceWeb#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.24.0/docs/resources/data_factory_linked_service_web#delete DataFactoryLinkedServiceWeb#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.24.0/docs/resources/data_factory_linked_service_web#read DataFactoryLinkedServiceWeb#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.24.0/docs/resources/data_factory_linked_service_web#update DataFactoryLinkedServiceWeb#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

