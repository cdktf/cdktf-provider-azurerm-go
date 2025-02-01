// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datafactorydatasetsnowflake


type DataFactoryDatasetSnowflakeTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.17.0/docs/resources/data_factory_dataset_snowflake#create DataFactoryDatasetSnowflake#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.17.0/docs/resources/data_factory_dataset_snowflake#delete DataFactoryDatasetSnowflake#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.17.0/docs/resources/data_factory_dataset_snowflake#read DataFactoryDatasetSnowflake#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.17.0/docs/resources/data_factory_dataset_snowflake#update DataFactoryDatasetSnowflake#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

