// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package streamanalyticsstreaminputiothub

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type StreamAnalyticsStreamInputIothubConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.82.0/docs/resources/stream_analytics_stream_input_iothub#endpoint StreamAnalyticsStreamInputIothub#endpoint}.
	Endpoint *string `field:"required" json:"endpoint" yaml:"endpoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.82.0/docs/resources/stream_analytics_stream_input_iothub#eventhub_consumer_group_name StreamAnalyticsStreamInputIothub#eventhub_consumer_group_name}.
	EventhubConsumerGroupName *string `field:"required" json:"eventhubConsumerGroupName" yaml:"eventhubConsumerGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.82.0/docs/resources/stream_analytics_stream_input_iothub#iothub_namespace StreamAnalyticsStreamInputIothub#iothub_namespace}.
	IothubNamespace *string `field:"required" json:"iothubNamespace" yaml:"iothubNamespace"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.82.0/docs/resources/stream_analytics_stream_input_iothub#name StreamAnalyticsStreamInputIothub#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.82.0/docs/resources/stream_analytics_stream_input_iothub#resource_group_name StreamAnalyticsStreamInputIothub#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// serialization block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.82.0/docs/resources/stream_analytics_stream_input_iothub#serialization StreamAnalyticsStreamInputIothub#serialization}
	Serialization *StreamAnalyticsStreamInputIothubSerialization `field:"required" json:"serialization" yaml:"serialization"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.82.0/docs/resources/stream_analytics_stream_input_iothub#shared_access_policy_key StreamAnalyticsStreamInputIothub#shared_access_policy_key}.
	SharedAccessPolicyKey *string `field:"required" json:"sharedAccessPolicyKey" yaml:"sharedAccessPolicyKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.82.0/docs/resources/stream_analytics_stream_input_iothub#shared_access_policy_name StreamAnalyticsStreamInputIothub#shared_access_policy_name}.
	SharedAccessPolicyName *string `field:"required" json:"sharedAccessPolicyName" yaml:"sharedAccessPolicyName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.82.0/docs/resources/stream_analytics_stream_input_iothub#stream_analytics_job_name StreamAnalyticsStreamInputIothub#stream_analytics_job_name}.
	StreamAnalyticsJobName *string `field:"required" json:"streamAnalyticsJobName" yaml:"streamAnalyticsJobName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.82.0/docs/resources/stream_analytics_stream_input_iothub#id StreamAnalyticsStreamInputIothub#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.82.0/docs/resources/stream_analytics_stream_input_iothub#timeouts StreamAnalyticsStreamInputIothub#timeouts}
	Timeouts *StreamAnalyticsStreamInputIothubTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

