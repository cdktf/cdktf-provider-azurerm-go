// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datafactorydatasetparquet


type DataFactoryDatasetParquetAzureBlobFsLocation struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.26.0/docs/resources/data_factory_dataset_parquet#dynamic_filename_enabled DataFactoryDatasetParquet#dynamic_filename_enabled}.
	DynamicFilenameEnabled interface{} `field:"optional" json:"dynamicFilenameEnabled" yaml:"dynamicFilenameEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.26.0/docs/resources/data_factory_dataset_parquet#dynamic_file_system_enabled DataFactoryDatasetParquet#dynamic_file_system_enabled}.
	DynamicFileSystemEnabled interface{} `field:"optional" json:"dynamicFileSystemEnabled" yaml:"dynamicFileSystemEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.26.0/docs/resources/data_factory_dataset_parquet#dynamic_path_enabled DataFactoryDatasetParquet#dynamic_path_enabled}.
	DynamicPathEnabled interface{} `field:"optional" json:"dynamicPathEnabled" yaml:"dynamicPathEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.26.0/docs/resources/data_factory_dataset_parquet#filename DataFactoryDatasetParquet#filename}.
	Filename *string `field:"optional" json:"filename" yaml:"filename"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.26.0/docs/resources/data_factory_dataset_parquet#file_system DataFactoryDatasetParquet#file_system}.
	FileSystem *string `field:"optional" json:"fileSystem" yaml:"fileSystem"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.26.0/docs/resources/data_factory_dataset_parquet#path DataFactoryDatasetParquet#path}.
	Path *string `field:"optional" json:"path" yaml:"path"`
}

