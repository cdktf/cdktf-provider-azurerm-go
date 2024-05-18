// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datafactorydatasetsqlservertable


type DataFactoryDatasetSqlServerTableSchemaColumn struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.0/docs/resources/data_factory_dataset_sql_server_table#name DataFactoryDatasetSqlServerTable#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.0/docs/resources/data_factory_dataset_sql_server_table#description DataFactoryDatasetSqlServerTable#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.0/docs/resources/data_factory_dataset_sql_server_table#type DataFactoryDatasetSqlServerTable#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

