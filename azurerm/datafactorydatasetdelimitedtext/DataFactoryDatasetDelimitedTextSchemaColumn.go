package datafactorydatasetdelimitedtext


type DataFactoryDatasetDelimitedTextSchemaColumn struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_dataset_delimited_text#name DataFactoryDatasetDelimitedText#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_dataset_delimited_text#description DataFactoryDatasetDelimitedText#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_dataset_delimited_text#type DataFactoryDatasetDelimitedText#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}
