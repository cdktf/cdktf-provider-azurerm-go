// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package iottimeseriesinsightsstandardenvironment

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IotTimeSeriesInsightsStandardEnvironmentConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.0/docs/resources/iot_time_series_insights_standard_environment#data_retention_time IotTimeSeriesInsightsStandardEnvironment#data_retention_time}.
	DataRetentionTime *string `field:"required" json:"dataRetentionTime" yaml:"dataRetentionTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.0/docs/resources/iot_time_series_insights_standard_environment#location IotTimeSeriesInsightsStandardEnvironment#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.0/docs/resources/iot_time_series_insights_standard_environment#name IotTimeSeriesInsightsStandardEnvironment#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.0/docs/resources/iot_time_series_insights_standard_environment#resource_group_name IotTimeSeriesInsightsStandardEnvironment#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.0/docs/resources/iot_time_series_insights_standard_environment#sku_name IotTimeSeriesInsightsStandardEnvironment#sku_name}.
	SkuName *string `field:"required" json:"skuName" yaml:"skuName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.0/docs/resources/iot_time_series_insights_standard_environment#id IotTimeSeriesInsightsStandardEnvironment#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.0/docs/resources/iot_time_series_insights_standard_environment#partition_key IotTimeSeriesInsightsStandardEnvironment#partition_key}.
	PartitionKey *string `field:"optional" json:"partitionKey" yaml:"partitionKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.0/docs/resources/iot_time_series_insights_standard_environment#storage_limit_exceeded_behavior IotTimeSeriesInsightsStandardEnvironment#storage_limit_exceeded_behavior}.
	StorageLimitExceededBehavior *string `field:"optional" json:"storageLimitExceededBehavior" yaml:"storageLimitExceededBehavior"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.0/docs/resources/iot_time_series_insights_standard_environment#tags IotTimeSeriesInsightsStandardEnvironment#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.0/docs/resources/iot_time_series_insights_standard_environment#timeouts IotTimeSeriesInsightsStandardEnvironment#timeouts}
	Timeouts *IotTimeSeriesInsightsStandardEnvironmentTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

