package datafactorydatasetmysql


type DataFactoryDatasetMysqlSchemaColumn struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_dataset_mysql#name DataFactoryDatasetMysql#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_dataset_mysql#description DataFactoryDatasetMysql#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_dataset_mysql#type DataFactoryDatasetMysql#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}
