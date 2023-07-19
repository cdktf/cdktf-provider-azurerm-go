package datafactorydatasetmysql


type DataFactoryDatasetMysqlTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/data_factory_dataset_mysql#create DataFactoryDatasetMysql#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/data_factory_dataset_mysql#delete DataFactoryDatasetMysql#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/data_factory_dataset_mysql#read DataFactoryDatasetMysql#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/data_factory_dataset_mysql#update DataFactoryDatasetMysql#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

