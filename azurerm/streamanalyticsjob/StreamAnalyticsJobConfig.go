// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package streamanalyticsjob

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type StreamAnalyticsJobConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.20.0/docs/resources/stream_analytics_job#location StreamAnalyticsJob#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.20.0/docs/resources/stream_analytics_job#name StreamAnalyticsJob#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.20.0/docs/resources/stream_analytics_job#resource_group_name StreamAnalyticsJob#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.20.0/docs/resources/stream_analytics_job#transformation_query StreamAnalyticsJob#transformation_query}.
	TransformationQuery *string `field:"required" json:"transformationQuery" yaml:"transformationQuery"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.20.0/docs/resources/stream_analytics_job#compatibility_level StreamAnalyticsJob#compatibility_level}.
	CompatibilityLevel *string `field:"optional" json:"compatibilityLevel" yaml:"compatibilityLevel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.20.0/docs/resources/stream_analytics_job#content_storage_policy StreamAnalyticsJob#content_storage_policy}.
	ContentStoragePolicy *string `field:"optional" json:"contentStoragePolicy" yaml:"contentStoragePolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.20.0/docs/resources/stream_analytics_job#data_locale StreamAnalyticsJob#data_locale}.
	DataLocale *string `field:"optional" json:"dataLocale" yaml:"dataLocale"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.20.0/docs/resources/stream_analytics_job#events_late_arrival_max_delay_in_seconds StreamAnalyticsJob#events_late_arrival_max_delay_in_seconds}.
	EventsLateArrivalMaxDelayInSeconds *float64 `field:"optional" json:"eventsLateArrivalMaxDelayInSeconds" yaml:"eventsLateArrivalMaxDelayInSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.20.0/docs/resources/stream_analytics_job#events_out_of_order_max_delay_in_seconds StreamAnalyticsJob#events_out_of_order_max_delay_in_seconds}.
	EventsOutOfOrderMaxDelayInSeconds *float64 `field:"optional" json:"eventsOutOfOrderMaxDelayInSeconds" yaml:"eventsOutOfOrderMaxDelayInSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.20.0/docs/resources/stream_analytics_job#events_out_of_order_policy StreamAnalyticsJob#events_out_of_order_policy}.
	EventsOutOfOrderPolicy *string `field:"optional" json:"eventsOutOfOrderPolicy" yaml:"eventsOutOfOrderPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.20.0/docs/resources/stream_analytics_job#id StreamAnalyticsJob#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// identity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.20.0/docs/resources/stream_analytics_job#identity StreamAnalyticsJob#identity}
	Identity *StreamAnalyticsJobIdentity `field:"optional" json:"identity" yaml:"identity"`
	// job_storage_account block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.20.0/docs/resources/stream_analytics_job#job_storage_account StreamAnalyticsJob#job_storage_account}
	JobStorageAccount interface{} `field:"optional" json:"jobStorageAccount" yaml:"jobStorageAccount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.20.0/docs/resources/stream_analytics_job#output_error_policy StreamAnalyticsJob#output_error_policy}.
	OutputErrorPolicy *string `field:"optional" json:"outputErrorPolicy" yaml:"outputErrorPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.20.0/docs/resources/stream_analytics_job#sku_name StreamAnalyticsJob#sku_name}.
	SkuName *string `field:"optional" json:"skuName" yaml:"skuName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.20.0/docs/resources/stream_analytics_job#stream_analytics_cluster_id StreamAnalyticsJob#stream_analytics_cluster_id}.
	StreamAnalyticsClusterId *string `field:"optional" json:"streamAnalyticsClusterId" yaml:"streamAnalyticsClusterId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.20.0/docs/resources/stream_analytics_job#streaming_units StreamAnalyticsJob#streaming_units}.
	StreamingUnits *float64 `field:"optional" json:"streamingUnits" yaml:"streamingUnits"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.20.0/docs/resources/stream_analytics_job#tags StreamAnalyticsJob#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.20.0/docs/resources/stream_analytics_job#timeouts StreamAnalyticsJob#timeouts}
	Timeouts *StreamAnalyticsJobTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.20.0/docs/resources/stream_analytics_job#type StreamAnalyticsJob#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

