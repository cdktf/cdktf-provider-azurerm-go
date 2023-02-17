package datafactorycustomdataset


type DataFactoryCustomDatasetTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_custom_dataset#create DataFactoryCustomDataset#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_custom_dataset#delete DataFactoryCustomDataset#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_custom_dataset#read DataFactoryCustomDataset#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_custom_dataset#update DataFactoryCustomDataset#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

