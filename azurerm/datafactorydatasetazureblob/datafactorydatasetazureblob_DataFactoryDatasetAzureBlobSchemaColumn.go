package datafactorydatasetazureblob


type DataFactoryDatasetAzureBlobSchemaColumn struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_dataset_azure_blob#name DataFactoryDatasetAzureBlob#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_dataset_azure_blob#description DataFactoryDatasetAzureBlob#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_dataset_azure_blob#type DataFactoryDatasetAzureBlob#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

