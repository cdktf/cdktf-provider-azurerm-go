// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datafactorydatasetsnowflake


type DataFactoryDatasetSnowflakeSchemaColumn struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.2.0/docs/resources/data_factory_dataset_snowflake#name DataFactoryDatasetSnowflake#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.2.0/docs/resources/data_factory_dataset_snowflake#precision DataFactoryDatasetSnowflake#precision}.
	Precision *float64 `field:"optional" json:"precision" yaml:"precision"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.2.0/docs/resources/data_factory_dataset_snowflake#scale DataFactoryDatasetSnowflake#scale}.
	Scale *float64 `field:"optional" json:"scale" yaml:"scale"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.2.0/docs/resources/data_factory_dataset_snowflake#type DataFactoryDatasetSnowflake#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

