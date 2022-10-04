package datafactorydatasethttp


type DataFactoryDatasetHttpSchemaColumn struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_dataset_http#name DataFactoryDatasetHttp#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_dataset_http#description DataFactoryDatasetHttp#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_dataset_http#type DataFactoryDatasetHttp#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

