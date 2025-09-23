// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datafactorydatasetazureblob

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataFactoryDatasetAzureBlobConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/data_factory_dataset_azure_blob#data_factory_id DataFactoryDatasetAzureBlob#data_factory_id}.
	DataFactoryId *string `field:"required" json:"dataFactoryId" yaml:"dataFactoryId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/data_factory_dataset_azure_blob#linked_service_name DataFactoryDatasetAzureBlob#linked_service_name}.
	LinkedServiceName *string `field:"required" json:"linkedServiceName" yaml:"linkedServiceName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/data_factory_dataset_azure_blob#name DataFactoryDatasetAzureBlob#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/data_factory_dataset_azure_blob#additional_properties DataFactoryDatasetAzureBlob#additional_properties}.
	AdditionalProperties *map[string]*string `field:"optional" json:"additionalProperties" yaml:"additionalProperties"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/data_factory_dataset_azure_blob#annotations DataFactoryDatasetAzureBlob#annotations}.
	Annotations *[]*string `field:"optional" json:"annotations" yaml:"annotations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/data_factory_dataset_azure_blob#description DataFactoryDatasetAzureBlob#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/data_factory_dataset_azure_blob#dynamic_filename_enabled DataFactoryDatasetAzureBlob#dynamic_filename_enabled}.
	DynamicFilenameEnabled interface{} `field:"optional" json:"dynamicFilenameEnabled" yaml:"dynamicFilenameEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/data_factory_dataset_azure_blob#dynamic_path_enabled DataFactoryDatasetAzureBlob#dynamic_path_enabled}.
	DynamicPathEnabled interface{} `field:"optional" json:"dynamicPathEnabled" yaml:"dynamicPathEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/data_factory_dataset_azure_blob#filename DataFactoryDatasetAzureBlob#filename}.
	Filename *string `field:"optional" json:"filename" yaml:"filename"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/data_factory_dataset_azure_blob#folder DataFactoryDatasetAzureBlob#folder}.
	Folder *string `field:"optional" json:"folder" yaml:"folder"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/data_factory_dataset_azure_blob#id DataFactoryDatasetAzureBlob#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/data_factory_dataset_azure_blob#parameters DataFactoryDatasetAzureBlob#parameters}.
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/data_factory_dataset_azure_blob#path DataFactoryDatasetAzureBlob#path}.
	Path *string `field:"optional" json:"path" yaml:"path"`
	// schema_column block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/data_factory_dataset_azure_blob#schema_column DataFactoryDatasetAzureBlob#schema_column}
	SchemaColumn interface{} `field:"optional" json:"schemaColumn" yaml:"schemaColumn"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/data_factory_dataset_azure_blob#timeouts DataFactoryDatasetAzureBlob#timeouts}
	Timeouts *DataFactoryDatasetAzureBlobTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

