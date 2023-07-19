package streamanalyticsstreaminputeventhubv2

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type StreamAnalyticsStreamInputEventhubV2Config struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/stream_analytics_stream_input_eventhub_v2#eventhub_name StreamAnalyticsStreamInputEventhubV2#eventhub_name}.
	EventhubName *string `field:"required" json:"eventhubName" yaml:"eventhubName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/stream_analytics_stream_input_eventhub_v2#name StreamAnalyticsStreamInputEventhubV2#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// serialization block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/stream_analytics_stream_input_eventhub_v2#serialization StreamAnalyticsStreamInputEventhubV2#serialization}
	Serialization *StreamAnalyticsStreamInputEventhubV2Serialization `field:"required" json:"serialization" yaml:"serialization"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/stream_analytics_stream_input_eventhub_v2#servicebus_namespace StreamAnalyticsStreamInputEventhubV2#servicebus_namespace}.
	ServicebusNamespace *string `field:"required" json:"servicebusNamespace" yaml:"servicebusNamespace"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/stream_analytics_stream_input_eventhub_v2#stream_analytics_job_id StreamAnalyticsStreamInputEventhubV2#stream_analytics_job_id}.
	StreamAnalyticsJobId *string `field:"required" json:"streamAnalyticsJobId" yaml:"streamAnalyticsJobId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/stream_analytics_stream_input_eventhub_v2#authentication_mode StreamAnalyticsStreamInputEventhubV2#authentication_mode}.
	AuthenticationMode *string `field:"optional" json:"authenticationMode" yaml:"authenticationMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/stream_analytics_stream_input_eventhub_v2#eventhub_consumer_group_name StreamAnalyticsStreamInputEventhubV2#eventhub_consumer_group_name}.
	EventhubConsumerGroupName *string `field:"optional" json:"eventhubConsumerGroupName" yaml:"eventhubConsumerGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/stream_analytics_stream_input_eventhub_v2#id StreamAnalyticsStreamInputEventhubV2#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/stream_analytics_stream_input_eventhub_v2#partition_key StreamAnalyticsStreamInputEventhubV2#partition_key}.
	PartitionKey *string `field:"optional" json:"partitionKey" yaml:"partitionKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/stream_analytics_stream_input_eventhub_v2#shared_access_policy_key StreamAnalyticsStreamInputEventhubV2#shared_access_policy_key}.
	SharedAccessPolicyKey *string `field:"optional" json:"sharedAccessPolicyKey" yaml:"sharedAccessPolicyKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/stream_analytics_stream_input_eventhub_v2#shared_access_policy_name StreamAnalyticsStreamInputEventhubV2#shared_access_policy_name}.
	SharedAccessPolicyName *string `field:"optional" json:"sharedAccessPolicyName" yaml:"sharedAccessPolicyName"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/stream_analytics_stream_input_eventhub_v2#timeouts StreamAnalyticsStreamInputEventhubV2#timeouts}
	Timeouts *StreamAnalyticsStreamInputEventhubV2Timeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

