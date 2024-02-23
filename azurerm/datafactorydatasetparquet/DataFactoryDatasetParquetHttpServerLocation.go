// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datafactorydatasetparquet


type DataFactoryDatasetParquetHttpServerLocation struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/data_factory_dataset_parquet#filename DataFactoryDatasetParquet#filename}.
	Filename *string `field:"required" json:"filename" yaml:"filename"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/data_factory_dataset_parquet#relative_url DataFactoryDatasetParquet#relative_url}.
	RelativeUrl *string `field:"required" json:"relativeUrl" yaml:"relativeUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/data_factory_dataset_parquet#dynamic_filename_enabled DataFactoryDatasetParquet#dynamic_filename_enabled}.
	DynamicFilenameEnabled interface{} `field:"optional" json:"dynamicFilenameEnabled" yaml:"dynamicFilenameEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/data_factory_dataset_parquet#dynamic_path_enabled DataFactoryDatasetParquet#dynamic_path_enabled}.
	DynamicPathEnabled interface{} `field:"optional" json:"dynamicPathEnabled" yaml:"dynamicPathEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/data_factory_dataset_parquet#path DataFactoryDatasetParquet#path}.
	Path *string `field:"optional" json:"path" yaml:"path"`
}

