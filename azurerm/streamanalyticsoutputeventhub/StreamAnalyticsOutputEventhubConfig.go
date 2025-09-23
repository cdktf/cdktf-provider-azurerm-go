// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package streamanalyticsoutputeventhub

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type StreamAnalyticsOutputEventhubConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/stream_analytics_output_eventhub#eventhub_name StreamAnalyticsOutputEventhub#eventhub_name}.
	EventhubName *string `field:"required" json:"eventhubName" yaml:"eventhubName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/stream_analytics_output_eventhub#name StreamAnalyticsOutputEventhub#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/stream_analytics_output_eventhub#resource_group_name StreamAnalyticsOutputEventhub#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// serialization block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/stream_analytics_output_eventhub#serialization StreamAnalyticsOutputEventhub#serialization}
	Serialization *StreamAnalyticsOutputEventhubSerialization `field:"required" json:"serialization" yaml:"serialization"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/stream_analytics_output_eventhub#servicebus_namespace StreamAnalyticsOutputEventhub#servicebus_namespace}.
	ServicebusNamespace *string `field:"required" json:"servicebusNamespace" yaml:"servicebusNamespace"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/stream_analytics_output_eventhub#stream_analytics_job_name StreamAnalyticsOutputEventhub#stream_analytics_job_name}.
	StreamAnalyticsJobName *string `field:"required" json:"streamAnalyticsJobName" yaml:"streamAnalyticsJobName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/stream_analytics_output_eventhub#authentication_mode StreamAnalyticsOutputEventhub#authentication_mode}.
	AuthenticationMode *string `field:"optional" json:"authenticationMode" yaml:"authenticationMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/stream_analytics_output_eventhub#id StreamAnalyticsOutputEventhub#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/stream_analytics_output_eventhub#partition_key StreamAnalyticsOutputEventhub#partition_key}.
	PartitionKey *string `field:"optional" json:"partitionKey" yaml:"partitionKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/stream_analytics_output_eventhub#property_columns StreamAnalyticsOutputEventhub#property_columns}.
	PropertyColumns *[]*string `field:"optional" json:"propertyColumns" yaml:"propertyColumns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/stream_analytics_output_eventhub#shared_access_policy_key StreamAnalyticsOutputEventhub#shared_access_policy_key}.
	SharedAccessPolicyKey *string `field:"optional" json:"sharedAccessPolicyKey" yaml:"sharedAccessPolicyKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/stream_analytics_output_eventhub#shared_access_policy_name StreamAnalyticsOutputEventhub#shared_access_policy_name}.
	SharedAccessPolicyName *string `field:"optional" json:"sharedAccessPolicyName" yaml:"sharedAccessPolicyName"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/stream_analytics_output_eventhub#timeouts StreamAnalyticsOutputEventhub#timeouts}
	Timeouts *StreamAnalyticsOutputEventhubTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

