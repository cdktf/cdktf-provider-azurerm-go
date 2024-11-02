// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datafactorydatasetpostgresql

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataFactoryDatasetPostgresqlConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/data_factory_dataset_postgresql#data_factory_id DataFactoryDatasetPostgresql#data_factory_id}.
	DataFactoryId *string `field:"required" json:"dataFactoryId" yaml:"dataFactoryId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/data_factory_dataset_postgresql#linked_service_name DataFactoryDatasetPostgresql#linked_service_name}.
	LinkedServiceName *string `field:"required" json:"linkedServiceName" yaml:"linkedServiceName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/data_factory_dataset_postgresql#name DataFactoryDatasetPostgresql#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/data_factory_dataset_postgresql#additional_properties DataFactoryDatasetPostgresql#additional_properties}.
	AdditionalProperties *map[string]*string `field:"optional" json:"additionalProperties" yaml:"additionalProperties"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/data_factory_dataset_postgresql#annotations DataFactoryDatasetPostgresql#annotations}.
	Annotations *[]*string `field:"optional" json:"annotations" yaml:"annotations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/data_factory_dataset_postgresql#description DataFactoryDatasetPostgresql#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/data_factory_dataset_postgresql#folder DataFactoryDatasetPostgresql#folder}.
	Folder *string `field:"optional" json:"folder" yaml:"folder"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/data_factory_dataset_postgresql#id DataFactoryDatasetPostgresql#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/data_factory_dataset_postgresql#parameters DataFactoryDatasetPostgresql#parameters}.
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
	// schema_column block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/data_factory_dataset_postgresql#schema_column DataFactoryDatasetPostgresql#schema_column}
	SchemaColumn interface{} `field:"optional" json:"schemaColumn" yaml:"schemaColumn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/data_factory_dataset_postgresql#table_name DataFactoryDatasetPostgresql#table_name}.
	TableName *string `field:"optional" json:"tableName" yaml:"tableName"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/data_factory_dataset_postgresql#timeouts DataFactoryDatasetPostgresql#timeouts}
	Timeouts *DataFactoryDatasetPostgresqlTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

