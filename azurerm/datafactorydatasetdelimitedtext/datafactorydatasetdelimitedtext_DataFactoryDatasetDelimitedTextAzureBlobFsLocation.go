package datafactorydatasetdelimitedtext


type DataFactoryDatasetDelimitedTextAzureBlobFsLocation struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_dataset_delimited_text#file_system DataFactoryDatasetDelimitedText#file_system}.
	FileSystem *string `field:"required" json:"fileSystem" yaml:"fileSystem"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_dataset_delimited_text#filename DataFactoryDatasetDelimitedText#filename}.
	Filename *string `field:"optional" json:"filename" yaml:"filename"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_dataset_delimited_text#path DataFactoryDatasetDelimitedText#path}.
	Path *string `field:"optional" json:"path" yaml:"path"`
}

