package datafactoryflowletdataflow


type DataFactoryFlowletDataFlowSourceFlowlet struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_flowlet_data_flow#name DataFactoryFlowletDataFlow#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_flowlet_data_flow#dataset_parameters DataFactoryFlowletDataFlow#dataset_parameters}.
	DatasetParameters *string `field:"optional" json:"datasetParameters" yaml:"datasetParameters"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_flowlet_data_flow#parameters DataFactoryFlowletDataFlow#parameters}.
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
}
