// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datafactoryflowletdataflow


type DataFactoryFlowletDataFlowSource struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.95.0/docs/resources/data_factory_flowlet_data_flow#name DataFactoryFlowletDataFlow#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// dataset block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.95.0/docs/resources/data_factory_flowlet_data_flow#dataset DataFactoryFlowletDataFlow#dataset}
	Dataset *DataFactoryFlowletDataFlowSourceDataset `field:"optional" json:"dataset" yaml:"dataset"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.95.0/docs/resources/data_factory_flowlet_data_flow#description DataFactoryFlowletDataFlow#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// flowlet block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.95.0/docs/resources/data_factory_flowlet_data_flow#flowlet DataFactoryFlowletDataFlow#flowlet}
	Flowlet *DataFactoryFlowletDataFlowSourceFlowlet `field:"optional" json:"flowlet" yaml:"flowlet"`
	// linked_service block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.95.0/docs/resources/data_factory_flowlet_data_flow#linked_service DataFactoryFlowletDataFlow#linked_service}
	LinkedService *DataFactoryFlowletDataFlowSourceLinkedService `field:"optional" json:"linkedService" yaml:"linkedService"`
	// rejected_linked_service block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.95.0/docs/resources/data_factory_flowlet_data_flow#rejected_linked_service DataFactoryFlowletDataFlow#rejected_linked_service}
	RejectedLinkedService *DataFactoryFlowletDataFlowSourceRejectedLinkedService `field:"optional" json:"rejectedLinkedService" yaml:"rejectedLinkedService"`
	// schema_linked_service block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.95.0/docs/resources/data_factory_flowlet_data_flow#schema_linked_service DataFactoryFlowletDataFlow#schema_linked_service}
	SchemaLinkedService *DataFactoryFlowletDataFlowSourceSchemaLinkedService `field:"optional" json:"schemaLinkedService" yaml:"schemaLinkedService"`
}

