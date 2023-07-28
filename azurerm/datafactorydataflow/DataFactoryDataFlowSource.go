package datafactorydataflow


type DataFactoryDataFlowSource struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/data_factory_data_flow#name DataFactoryDataFlow#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// dataset block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/data_factory_data_flow#dataset DataFactoryDataFlow#dataset}
	Dataset *DataFactoryDataFlowSourceDataset `field:"optional" json:"dataset" yaml:"dataset"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/data_factory_data_flow#description DataFactoryDataFlow#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// flowlet block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/data_factory_data_flow#flowlet DataFactoryDataFlow#flowlet}
	Flowlet *DataFactoryDataFlowSourceFlowlet `field:"optional" json:"flowlet" yaml:"flowlet"`
	// linked_service block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/data_factory_data_flow#linked_service DataFactoryDataFlow#linked_service}
	LinkedService *DataFactoryDataFlowSourceLinkedService `field:"optional" json:"linkedService" yaml:"linkedService"`
	// rejected_linked_service block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/data_factory_data_flow#rejected_linked_service DataFactoryDataFlow#rejected_linked_service}
	RejectedLinkedService *DataFactoryDataFlowSourceRejectedLinkedService `field:"optional" json:"rejectedLinkedService" yaml:"rejectedLinkedService"`
	// schema_linked_service block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/data_factory_data_flow#schema_linked_service DataFactoryDataFlow#schema_linked_service}
	SchemaLinkedService *DataFactoryDataFlowSourceSchemaLinkedService `field:"optional" json:"schemaLinkedService" yaml:"schemaLinkedService"`
}

