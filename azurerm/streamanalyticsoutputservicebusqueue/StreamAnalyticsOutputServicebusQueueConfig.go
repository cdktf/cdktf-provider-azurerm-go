// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package streamanalyticsoutputservicebusqueue

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type StreamAnalyticsOutputServicebusQueueConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/stream_analytics_output_servicebus_queue#name StreamAnalyticsOutputServicebusQueue#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/stream_analytics_output_servicebus_queue#queue_name StreamAnalyticsOutputServicebusQueue#queue_name}.
	QueueName *string `field:"required" json:"queueName" yaml:"queueName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/stream_analytics_output_servicebus_queue#resource_group_name StreamAnalyticsOutputServicebusQueue#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// serialization block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/stream_analytics_output_servicebus_queue#serialization StreamAnalyticsOutputServicebusQueue#serialization}
	Serialization *StreamAnalyticsOutputServicebusQueueSerialization `field:"required" json:"serialization" yaml:"serialization"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/stream_analytics_output_servicebus_queue#servicebus_namespace StreamAnalyticsOutputServicebusQueue#servicebus_namespace}.
	ServicebusNamespace *string `field:"required" json:"servicebusNamespace" yaml:"servicebusNamespace"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/stream_analytics_output_servicebus_queue#stream_analytics_job_name StreamAnalyticsOutputServicebusQueue#stream_analytics_job_name}.
	StreamAnalyticsJobName *string `field:"required" json:"streamAnalyticsJobName" yaml:"streamAnalyticsJobName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/stream_analytics_output_servicebus_queue#authentication_mode StreamAnalyticsOutputServicebusQueue#authentication_mode}.
	AuthenticationMode *string `field:"optional" json:"authenticationMode" yaml:"authenticationMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/stream_analytics_output_servicebus_queue#id StreamAnalyticsOutputServicebusQueue#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/stream_analytics_output_servicebus_queue#property_columns StreamAnalyticsOutputServicebusQueue#property_columns}.
	PropertyColumns *[]*string `field:"optional" json:"propertyColumns" yaml:"propertyColumns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/stream_analytics_output_servicebus_queue#shared_access_policy_key StreamAnalyticsOutputServicebusQueue#shared_access_policy_key}.
	SharedAccessPolicyKey *string `field:"optional" json:"sharedAccessPolicyKey" yaml:"sharedAccessPolicyKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/stream_analytics_output_servicebus_queue#shared_access_policy_name StreamAnalyticsOutputServicebusQueue#shared_access_policy_name}.
	SharedAccessPolicyName *string `field:"optional" json:"sharedAccessPolicyName" yaml:"sharedAccessPolicyName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/stream_analytics_output_servicebus_queue#system_property_columns StreamAnalyticsOutputServicebusQueue#system_property_columns}.
	SystemPropertyColumns *map[string]*string `field:"optional" json:"systemPropertyColumns" yaml:"systemPropertyColumns"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/stream_analytics_output_servicebus_queue#timeouts StreamAnalyticsOutputServicebusQueue#timeouts}
	Timeouts *StreamAnalyticsOutputServicebusQueueTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

