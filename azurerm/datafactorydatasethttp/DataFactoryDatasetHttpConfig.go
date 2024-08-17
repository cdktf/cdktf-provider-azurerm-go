// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datafactorydatasethttp

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataFactoryDatasetHttpConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/data_factory_dataset_http#data_factory_id DataFactoryDatasetHttp#data_factory_id}.
	DataFactoryId *string `field:"required" json:"dataFactoryId" yaml:"dataFactoryId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/data_factory_dataset_http#linked_service_name DataFactoryDatasetHttp#linked_service_name}.
	LinkedServiceName *string `field:"required" json:"linkedServiceName" yaml:"linkedServiceName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/data_factory_dataset_http#name DataFactoryDatasetHttp#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/data_factory_dataset_http#additional_properties DataFactoryDatasetHttp#additional_properties}.
	AdditionalProperties *map[string]*string `field:"optional" json:"additionalProperties" yaml:"additionalProperties"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/data_factory_dataset_http#annotations DataFactoryDatasetHttp#annotations}.
	Annotations *[]*string `field:"optional" json:"annotations" yaml:"annotations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/data_factory_dataset_http#description DataFactoryDatasetHttp#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/data_factory_dataset_http#folder DataFactoryDatasetHttp#folder}.
	Folder *string `field:"optional" json:"folder" yaml:"folder"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/data_factory_dataset_http#id DataFactoryDatasetHttp#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/data_factory_dataset_http#parameters DataFactoryDatasetHttp#parameters}.
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/data_factory_dataset_http#relative_url DataFactoryDatasetHttp#relative_url}.
	RelativeUrl *string `field:"optional" json:"relativeUrl" yaml:"relativeUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/data_factory_dataset_http#request_body DataFactoryDatasetHttp#request_body}.
	RequestBody *string `field:"optional" json:"requestBody" yaml:"requestBody"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/data_factory_dataset_http#request_method DataFactoryDatasetHttp#request_method}.
	RequestMethod *string `field:"optional" json:"requestMethod" yaml:"requestMethod"`
	// schema_column block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/data_factory_dataset_http#schema_column DataFactoryDatasetHttp#schema_column}
	SchemaColumn interface{} `field:"optional" json:"schemaColumn" yaml:"schemaColumn"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/data_factory_dataset_http#timeouts DataFactoryDatasetHttp#timeouts}
	Timeouts *DataFactoryDatasetHttpTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

