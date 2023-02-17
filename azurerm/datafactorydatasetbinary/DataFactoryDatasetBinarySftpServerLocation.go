package datafactorydatasetbinary


type DataFactoryDatasetBinarySftpServerLocation struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_dataset_binary#filename DataFactoryDatasetBinary#filename}.
	Filename *string `field:"required" json:"filename" yaml:"filename"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_dataset_binary#path DataFactoryDatasetBinary#path}.
	Path *string `field:"required" json:"path" yaml:"path"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_dataset_binary#dynamic_filename_enabled DataFactoryDatasetBinary#dynamic_filename_enabled}.
	DynamicFilenameEnabled interface{} `field:"optional" json:"dynamicFilenameEnabled" yaml:"dynamicFilenameEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_dataset_binary#dynamic_path_enabled DataFactoryDatasetBinary#dynamic_path_enabled}.
	DynamicPathEnabled interface{} `field:"optional" json:"dynamicPathEnabled" yaml:"dynamicPathEnabled"`
}

