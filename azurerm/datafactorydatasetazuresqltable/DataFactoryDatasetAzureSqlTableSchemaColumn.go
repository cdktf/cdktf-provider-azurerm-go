// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datafactorydatasetazuresqltable


type DataFactoryDatasetAzureSqlTableSchemaColumn struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/data_factory_dataset_azure_sql_table#name DataFactoryDatasetAzureSqlTable#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/data_factory_dataset_azure_sql_table#description DataFactoryDatasetAzureSqlTable#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/data_factory_dataset_azure_sql_table#type DataFactoryDatasetAzureSqlTable#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

