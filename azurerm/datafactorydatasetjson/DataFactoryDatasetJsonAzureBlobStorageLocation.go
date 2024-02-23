// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datafactorydatasetjson


type DataFactoryDatasetJsonAzureBlobStorageLocation struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/data_factory_dataset_json#container DataFactoryDatasetJson#container}.
	Container *string `field:"required" json:"container" yaml:"container"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/data_factory_dataset_json#filename DataFactoryDatasetJson#filename}.
	Filename *string `field:"required" json:"filename" yaml:"filename"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/data_factory_dataset_json#path DataFactoryDatasetJson#path}.
	Path *string `field:"required" json:"path" yaml:"path"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/data_factory_dataset_json#dynamic_container_enabled DataFactoryDatasetJson#dynamic_container_enabled}.
	DynamicContainerEnabled interface{} `field:"optional" json:"dynamicContainerEnabled" yaml:"dynamicContainerEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/data_factory_dataset_json#dynamic_filename_enabled DataFactoryDatasetJson#dynamic_filename_enabled}.
	DynamicFilenameEnabled interface{} `field:"optional" json:"dynamicFilenameEnabled" yaml:"dynamicFilenameEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/data_factory_dataset_json#dynamic_path_enabled DataFactoryDatasetJson#dynamic_path_enabled}.
	DynamicPathEnabled interface{} `field:"optional" json:"dynamicPathEnabled" yaml:"dynamicPathEnabled"`
}

