package datafactorydatasetpostgresql


type DataFactoryDatasetPostgresqlTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/data_factory_dataset_postgresql#create DataFactoryDatasetPostgresql#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/data_factory_dataset_postgresql#delete DataFactoryDatasetPostgresql#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/data_factory_dataset_postgresql#read DataFactoryDatasetPostgresql#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/data_factory_dataset_postgresql#update DataFactoryDatasetPostgresql#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

