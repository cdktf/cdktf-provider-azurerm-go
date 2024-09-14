// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datafactorydatasetparquet

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataFactoryDatasetParquetConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.2.0/docs/resources/data_factory_dataset_parquet#data_factory_id DataFactoryDatasetParquet#data_factory_id}.
	DataFactoryId *string `field:"required" json:"dataFactoryId" yaml:"dataFactoryId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.2.0/docs/resources/data_factory_dataset_parquet#linked_service_name DataFactoryDatasetParquet#linked_service_name}.
	LinkedServiceName *string `field:"required" json:"linkedServiceName" yaml:"linkedServiceName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.2.0/docs/resources/data_factory_dataset_parquet#name DataFactoryDatasetParquet#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.2.0/docs/resources/data_factory_dataset_parquet#additional_properties DataFactoryDatasetParquet#additional_properties}.
	AdditionalProperties *map[string]*string `field:"optional" json:"additionalProperties" yaml:"additionalProperties"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.2.0/docs/resources/data_factory_dataset_parquet#annotations DataFactoryDatasetParquet#annotations}.
	Annotations *[]*string `field:"optional" json:"annotations" yaml:"annotations"`
	// azure_blob_fs_location block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.2.0/docs/resources/data_factory_dataset_parquet#azure_blob_fs_location DataFactoryDatasetParquet#azure_blob_fs_location}
	AzureBlobFsLocation *DataFactoryDatasetParquetAzureBlobFsLocation `field:"optional" json:"azureBlobFsLocation" yaml:"azureBlobFsLocation"`
	// azure_blob_storage_location block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.2.0/docs/resources/data_factory_dataset_parquet#azure_blob_storage_location DataFactoryDatasetParquet#azure_blob_storage_location}
	AzureBlobStorageLocation *DataFactoryDatasetParquetAzureBlobStorageLocation `field:"optional" json:"azureBlobStorageLocation" yaml:"azureBlobStorageLocation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.2.0/docs/resources/data_factory_dataset_parquet#compression_codec DataFactoryDatasetParquet#compression_codec}.
	CompressionCodec *string `field:"optional" json:"compressionCodec" yaml:"compressionCodec"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.2.0/docs/resources/data_factory_dataset_parquet#compression_level DataFactoryDatasetParquet#compression_level}.
	CompressionLevel *string `field:"optional" json:"compressionLevel" yaml:"compressionLevel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.2.0/docs/resources/data_factory_dataset_parquet#description DataFactoryDatasetParquet#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.2.0/docs/resources/data_factory_dataset_parquet#folder DataFactoryDatasetParquet#folder}.
	Folder *string `field:"optional" json:"folder" yaml:"folder"`
	// http_server_location block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.2.0/docs/resources/data_factory_dataset_parquet#http_server_location DataFactoryDatasetParquet#http_server_location}
	HttpServerLocation *DataFactoryDatasetParquetHttpServerLocation `field:"optional" json:"httpServerLocation" yaml:"httpServerLocation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.2.0/docs/resources/data_factory_dataset_parquet#id DataFactoryDatasetParquet#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.2.0/docs/resources/data_factory_dataset_parquet#parameters DataFactoryDatasetParquet#parameters}.
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
	// schema_column block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.2.0/docs/resources/data_factory_dataset_parquet#schema_column DataFactoryDatasetParquet#schema_column}
	SchemaColumn interface{} `field:"optional" json:"schemaColumn" yaml:"schemaColumn"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.2.0/docs/resources/data_factory_dataset_parquet#timeouts DataFactoryDatasetParquet#timeouts}
	Timeouts *DataFactoryDatasetParquetTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

