// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package streamanalyticsstreaminputblob

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type StreamAnalyticsStreamInputBlobConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.79.0/docs/resources/stream_analytics_stream_input_blob#date_format StreamAnalyticsStreamInputBlob#date_format}.
	DateFormat *string `field:"required" json:"dateFormat" yaml:"dateFormat"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.79.0/docs/resources/stream_analytics_stream_input_blob#name StreamAnalyticsStreamInputBlob#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.79.0/docs/resources/stream_analytics_stream_input_blob#path_pattern StreamAnalyticsStreamInputBlob#path_pattern}.
	PathPattern *string `field:"required" json:"pathPattern" yaml:"pathPattern"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.79.0/docs/resources/stream_analytics_stream_input_blob#resource_group_name StreamAnalyticsStreamInputBlob#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// serialization block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.79.0/docs/resources/stream_analytics_stream_input_blob#serialization StreamAnalyticsStreamInputBlob#serialization}
	Serialization *StreamAnalyticsStreamInputBlobSerialization `field:"required" json:"serialization" yaml:"serialization"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.79.0/docs/resources/stream_analytics_stream_input_blob#storage_account_key StreamAnalyticsStreamInputBlob#storage_account_key}.
	StorageAccountKey *string `field:"required" json:"storageAccountKey" yaml:"storageAccountKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.79.0/docs/resources/stream_analytics_stream_input_blob#storage_account_name StreamAnalyticsStreamInputBlob#storage_account_name}.
	StorageAccountName *string `field:"required" json:"storageAccountName" yaml:"storageAccountName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.79.0/docs/resources/stream_analytics_stream_input_blob#storage_container_name StreamAnalyticsStreamInputBlob#storage_container_name}.
	StorageContainerName *string `field:"required" json:"storageContainerName" yaml:"storageContainerName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.79.0/docs/resources/stream_analytics_stream_input_blob#stream_analytics_job_name StreamAnalyticsStreamInputBlob#stream_analytics_job_name}.
	StreamAnalyticsJobName *string `field:"required" json:"streamAnalyticsJobName" yaml:"streamAnalyticsJobName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.79.0/docs/resources/stream_analytics_stream_input_blob#time_format StreamAnalyticsStreamInputBlob#time_format}.
	TimeFormat *string `field:"required" json:"timeFormat" yaml:"timeFormat"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.79.0/docs/resources/stream_analytics_stream_input_blob#id StreamAnalyticsStreamInputBlob#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.79.0/docs/resources/stream_analytics_stream_input_blob#timeouts StreamAnalyticsStreamInputBlob#timeouts}
	Timeouts *StreamAnalyticsStreamInputBlobTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

