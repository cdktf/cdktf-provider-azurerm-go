package datafactorydataflow


type DataFactoryDataFlowTransformationFlowlet struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/data_factory_data_flow#name DataFactoryDataFlow#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/data_factory_data_flow#dataset_parameters DataFactoryDataFlow#dataset_parameters}.
	DatasetParameters *string `field:"optional" json:"datasetParameters" yaml:"datasetParameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/data_factory_data_flow#parameters DataFactoryDataFlow#parameters}.
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
}

