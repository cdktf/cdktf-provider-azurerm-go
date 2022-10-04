package datafactorydatasetsqlservertable


type DataFactoryDatasetSqlServerTableSchemaColumn struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_dataset_sql_server_table#name DataFactoryDatasetSqlServerTable#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_dataset_sql_server_table#description DataFactoryDatasetSqlServerTable#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_dataset_sql_server_table#type DataFactoryDatasetSqlServerTable#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

