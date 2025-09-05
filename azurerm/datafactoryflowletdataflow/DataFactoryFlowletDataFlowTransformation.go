// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datafactoryflowletdataflow


type DataFactoryFlowletDataFlowTransformation struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/data_factory_flowlet_data_flow#name DataFactoryFlowletDataFlow#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// dataset block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/data_factory_flowlet_data_flow#dataset DataFactoryFlowletDataFlow#dataset}
	Dataset *DataFactoryFlowletDataFlowTransformationDataset `field:"optional" json:"dataset" yaml:"dataset"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/data_factory_flowlet_data_flow#description DataFactoryFlowletDataFlow#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// flowlet block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/data_factory_flowlet_data_flow#flowlet DataFactoryFlowletDataFlow#flowlet}
	Flowlet *DataFactoryFlowletDataFlowTransformationFlowlet `field:"optional" json:"flowlet" yaml:"flowlet"`
	// linked_service block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/data_factory_flowlet_data_flow#linked_service DataFactoryFlowletDataFlow#linked_service}
	LinkedService *DataFactoryFlowletDataFlowTransformationLinkedService `field:"optional" json:"linkedService" yaml:"linkedService"`
}

