package datafactoryflowletdataflow


type DataFactoryFlowletDataFlowTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_flowlet_data_flow#create DataFactoryFlowletDataFlow#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_flowlet_data_flow#delete DataFactoryFlowletDataFlow#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_flowlet_data_flow#read DataFactoryFlowletDataFlow#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_flowlet_data_flow#update DataFactoryFlowletDataFlow#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}
