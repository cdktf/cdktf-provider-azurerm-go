// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datafactorydatasetsqlservertable


type DataFactoryDatasetSqlServerTableTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.0/docs/resources/data_factory_dataset_sql_server_table#create DataFactoryDatasetSqlServerTable#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.0/docs/resources/data_factory_dataset_sql_server_table#delete DataFactoryDatasetSqlServerTable#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.0/docs/resources/data_factory_dataset_sql_server_table#read DataFactoryDatasetSqlServerTable#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.0/docs/resources/data_factory_dataset_sql_server_table#update DataFactoryDatasetSqlServerTable#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

