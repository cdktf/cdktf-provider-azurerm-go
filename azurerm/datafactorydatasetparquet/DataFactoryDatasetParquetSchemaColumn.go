// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datafactorydatasetparquet


type DataFactoryDatasetParquetSchemaColumn struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.97.1/docs/resources/data_factory_dataset_parquet#name DataFactoryDatasetParquet#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.97.1/docs/resources/data_factory_dataset_parquet#description DataFactoryDatasetParquet#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.97.1/docs/resources/data_factory_dataset_parquet#type DataFactoryDatasetParquet#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

