// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datafactorydatasetazuresqltable

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataFactoryDatasetAzureSqlTableConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/data_factory_dataset_azure_sql_table#data_factory_id DataFactoryDatasetAzureSqlTable#data_factory_id}.
	DataFactoryId *string `field:"required" json:"dataFactoryId" yaml:"dataFactoryId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/data_factory_dataset_azure_sql_table#linked_service_id DataFactoryDatasetAzureSqlTable#linked_service_id}.
	LinkedServiceId *string `field:"required" json:"linkedServiceId" yaml:"linkedServiceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/data_factory_dataset_azure_sql_table#name DataFactoryDatasetAzureSqlTable#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/data_factory_dataset_azure_sql_table#additional_properties DataFactoryDatasetAzureSqlTable#additional_properties}.
	AdditionalProperties *map[string]*string `field:"optional" json:"additionalProperties" yaml:"additionalProperties"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/data_factory_dataset_azure_sql_table#annotations DataFactoryDatasetAzureSqlTable#annotations}.
	Annotations *[]*string `field:"optional" json:"annotations" yaml:"annotations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/data_factory_dataset_azure_sql_table#description DataFactoryDatasetAzureSqlTable#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/data_factory_dataset_azure_sql_table#folder DataFactoryDatasetAzureSqlTable#folder}.
	Folder *string `field:"optional" json:"folder" yaml:"folder"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/data_factory_dataset_azure_sql_table#id DataFactoryDatasetAzureSqlTable#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/data_factory_dataset_azure_sql_table#parameters DataFactoryDatasetAzureSqlTable#parameters}.
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/data_factory_dataset_azure_sql_table#schema DataFactoryDatasetAzureSqlTable#schema}.
	Schema *string `field:"optional" json:"schema" yaml:"schema"`
	// schema_column block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/data_factory_dataset_azure_sql_table#schema_column DataFactoryDatasetAzureSqlTable#schema_column}
	SchemaColumn interface{} `field:"optional" json:"schemaColumn" yaml:"schemaColumn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/data_factory_dataset_azure_sql_table#table DataFactoryDatasetAzureSqlTable#table}.
	Table *string `field:"optional" json:"table" yaml:"table"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/data_factory_dataset_azure_sql_table#timeouts DataFactoryDatasetAzureSqlTable#timeouts}
	Timeouts *DataFactoryDatasetAzureSqlTableTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

