package datafactorydatasetjson


type DataFactoryDatasetJsonHttpServerLocation struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.56.0/docs/resources/data_factory_dataset_json#filename DataFactoryDatasetJson#filename}.
	Filename *string `field:"required" json:"filename" yaml:"filename"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.56.0/docs/resources/data_factory_dataset_json#path DataFactoryDatasetJson#path}.
	Path *string `field:"required" json:"path" yaml:"path"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.56.0/docs/resources/data_factory_dataset_json#relative_url DataFactoryDatasetJson#relative_url}.
	RelativeUrl *string `field:"required" json:"relativeUrl" yaml:"relativeUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.56.0/docs/resources/data_factory_dataset_json#dynamic_filename_enabled DataFactoryDatasetJson#dynamic_filename_enabled}.
	DynamicFilenameEnabled interface{} `field:"optional" json:"dynamicFilenameEnabled" yaml:"dynamicFilenameEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.56.0/docs/resources/data_factory_dataset_json#dynamic_path_enabled DataFactoryDatasetJson#dynamic_path_enabled}.
	DynamicPathEnabled interface{} `field:"optional" json:"dynamicPathEnabled" yaml:"dynamicPathEnabled"`
}

