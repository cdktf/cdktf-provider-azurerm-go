// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package streamanalyticsoutputblob

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type StreamAnalyticsOutputBlobConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.28.0/docs/resources/stream_analytics_output_blob#date_format StreamAnalyticsOutputBlob#date_format}.
	DateFormat *string `field:"required" json:"dateFormat" yaml:"dateFormat"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.28.0/docs/resources/stream_analytics_output_blob#name StreamAnalyticsOutputBlob#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.28.0/docs/resources/stream_analytics_output_blob#path_pattern StreamAnalyticsOutputBlob#path_pattern}.
	PathPattern *string `field:"required" json:"pathPattern" yaml:"pathPattern"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.28.0/docs/resources/stream_analytics_output_blob#resource_group_name StreamAnalyticsOutputBlob#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// serialization block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.28.0/docs/resources/stream_analytics_output_blob#serialization StreamAnalyticsOutputBlob#serialization}
	Serialization *StreamAnalyticsOutputBlobSerialization `field:"required" json:"serialization" yaml:"serialization"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.28.0/docs/resources/stream_analytics_output_blob#storage_account_name StreamAnalyticsOutputBlob#storage_account_name}.
	StorageAccountName *string `field:"required" json:"storageAccountName" yaml:"storageAccountName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.28.0/docs/resources/stream_analytics_output_blob#storage_container_name StreamAnalyticsOutputBlob#storage_container_name}.
	StorageContainerName *string `field:"required" json:"storageContainerName" yaml:"storageContainerName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.28.0/docs/resources/stream_analytics_output_blob#stream_analytics_job_name StreamAnalyticsOutputBlob#stream_analytics_job_name}.
	StreamAnalyticsJobName *string `field:"required" json:"streamAnalyticsJobName" yaml:"streamAnalyticsJobName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.28.0/docs/resources/stream_analytics_output_blob#time_format StreamAnalyticsOutputBlob#time_format}.
	TimeFormat *string `field:"required" json:"timeFormat" yaml:"timeFormat"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.28.0/docs/resources/stream_analytics_output_blob#authentication_mode StreamAnalyticsOutputBlob#authentication_mode}.
	AuthenticationMode *string `field:"optional" json:"authenticationMode" yaml:"authenticationMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.28.0/docs/resources/stream_analytics_output_blob#batch_max_wait_time StreamAnalyticsOutputBlob#batch_max_wait_time}.
	BatchMaxWaitTime *string `field:"optional" json:"batchMaxWaitTime" yaml:"batchMaxWaitTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.28.0/docs/resources/stream_analytics_output_blob#batch_min_rows StreamAnalyticsOutputBlob#batch_min_rows}.
	BatchMinRows *float64 `field:"optional" json:"batchMinRows" yaml:"batchMinRows"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.28.0/docs/resources/stream_analytics_output_blob#blob_write_mode StreamAnalyticsOutputBlob#blob_write_mode}.
	BlobWriteMode *string `field:"optional" json:"blobWriteMode" yaml:"blobWriteMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.28.0/docs/resources/stream_analytics_output_blob#id StreamAnalyticsOutputBlob#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.28.0/docs/resources/stream_analytics_output_blob#storage_account_key StreamAnalyticsOutputBlob#storage_account_key}.
	StorageAccountKey *string `field:"optional" json:"storageAccountKey" yaml:"storageAccountKey"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.28.0/docs/resources/stream_analytics_output_blob#timeouts StreamAnalyticsOutputBlob#timeouts}
	Timeouts *StreamAnalyticsOutputBlobTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

