// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package eventgridsystemtopiceventsubscription

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EventgridSystemTopicEventSubscriptionConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/eventgrid_system_topic_event_subscription#name EventgridSystemTopicEventSubscription#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/eventgrid_system_topic_event_subscription#resource_group_name EventgridSystemTopicEventSubscription#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/eventgrid_system_topic_event_subscription#system_topic EventgridSystemTopicEventSubscription#system_topic}.
	SystemTopic *string `field:"required" json:"systemTopic" yaml:"systemTopic"`
	// advanced_filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/eventgrid_system_topic_event_subscription#advanced_filter EventgridSystemTopicEventSubscription#advanced_filter}
	AdvancedFilter *EventgridSystemTopicEventSubscriptionAdvancedFilter `field:"optional" json:"advancedFilter" yaml:"advancedFilter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/eventgrid_system_topic_event_subscription#advanced_filtering_on_arrays_enabled EventgridSystemTopicEventSubscription#advanced_filtering_on_arrays_enabled}.
	AdvancedFilteringOnArraysEnabled interface{} `field:"optional" json:"advancedFilteringOnArraysEnabled" yaml:"advancedFilteringOnArraysEnabled"`
	// azure_function_endpoint block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/eventgrid_system_topic_event_subscription#azure_function_endpoint EventgridSystemTopicEventSubscription#azure_function_endpoint}
	AzureFunctionEndpoint *EventgridSystemTopicEventSubscriptionAzureFunctionEndpoint `field:"optional" json:"azureFunctionEndpoint" yaml:"azureFunctionEndpoint"`
	// dead_letter_identity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/eventgrid_system_topic_event_subscription#dead_letter_identity EventgridSystemTopicEventSubscription#dead_letter_identity}
	DeadLetterIdentity *EventgridSystemTopicEventSubscriptionDeadLetterIdentity `field:"optional" json:"deadLetterIdentity" yaml:"deadLetterIdentity"`
	// delivery_identity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/eventgrid_system_topic_event_subscription#delivery_identity EventgridSystemTopicEventSubscription#delivery_identity}
	DeliveryIdentity *EventgridSystemTopicEventSubscriptionDeliveryIdentity `field:"optional" json:"deliveryIdentity" yaml:"deliveryIdentity"`
	// delivery_property block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/eventgrid_system_topic_event_subscription#delivery_property EventgridSystemTopicEventSubscription#delivery_property}
	DeliveryProperty interface{} `field:"optional" json:"deliveryProperty" yaml:"deliveryProperty"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/eventgrid_system_topic_event_subscription#event_delivery_schema EventgridSystemTopicEventSubscription#event_delivery_schema}.
	EventDeliverySchema *string `field:"optional" json:"eventDeliverySchema" yaml:"eventDeliverySchema"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/eventgrid_system_topic_event_subscription#eventhub_endpoint_id EventgridSystemTopicEventSubscription#eventhub_endpoint_id}.
	EventhubEndpointId *string `field:"optional" json:"eventhubEndpointId" yaml:"eventhubEndpointId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/eventgrid_system_topic_event_subscription#expiration_time_utc EventgridSystemTopicEventSubscription#expiration_time_utc}.
	ExpirationTimeUtc *string `field:"optional" json:"expirationTimeUtc" yaml:"expirationTimeUtc"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/eventgrid_system_topic_event_subscription#hybrid_connection_endpoint_id EventgridSystemTopicEventSubscription#hybrid_connection_endpoint_id}.
	HybridConnectionEndpointId *string `field:"optional" json:"hybridConnectionEndpointId" yaml:"hybridConnectionEndpointId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/eventgrid_system_topic_event_subscription#id EventgridSystemTopicEventSubscription#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/eventgrid_system_topic_event_subscription#included_event_types EventgridSystemTopicEventSubscription#included_event_types}.
	IncludedEventTypes *[]*string `field:"optional" json:"includedEventTypes" yaml:"includedEventTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/eventgrid_system_topic_event_subscription#labels EventgridSystemTopicEventSubscription#labels}.
	Labels *[]*string `field:"optional" json:"labels" yaml:"labels"`
	// retry_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/eventgrid_system_topic_event_subscription#retry_policy EventgridSystemTopicEventSubscription#retry_policy}
	RetryPolicy *EventgridSystemTopicEventSubscriptionRetryPolicy `field:"optional" json:"retryPolicy" yaml:"retryPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/eventgrid_system_topic_event_subscription#service_bus_queue_endpoint_id EventgridSystemTopicEventSubscription#service_bus_queue_endpoint_id}.
	ServiceBusQueueEndpointId *string `field:"optional" json:"serviceBusQueueEndpointId" yaml:"serviceBusQueueEndpointId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/eventgrid_system_topic_event_subscription#service_bus_topic_endpoint_id EventgridSystemTopicEventSubscription#service_bus_topic_endpoint_id}.
	ServiceBusTopicEndpointId *string `field:"optional" json:"serviceBusTopicEndpointId" yaml:"serviceBusTopicEndpointId"`
	// storage_blob_dead_letter_destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/eventgrid_system_topic_event_subscription#storage_blob_dead_letter_destination EventgridSystemTopicEventSubscription#storage_blob_dead_letter_destination}
	StorageBlobDeadLetterDestination *EventgridSystemTopicEventSubscriptionStorageBlobDeadLetterDestination `field:"optional" json:"storageBlobDeadLetterDestination" yaml:"storageBlobDeadLetterDestination"`
	// storage_queue_endpoint block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/eventgrid_system_topic_event_subscription#storage_queue_endpoint EventgridSystemTopicEventSubscription#storage_queue_endpoint}
	StorageQueueEndpoint *EventgridSystemTopicEventSubscriptionStorageQueueEndpoint `field:"optional" json:"storageQueueEndpoint" yaml:"storageQueueEndpoint"`
	// subject_filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/eventgrid_system_topic_event_subscription#subject_filter EventgridSystemTopicEventSubscription#subject_filter}
	SubjectFilter *EventgridSystemTopicEventSubscriptionSubjectFilter `field:"optional" json:"subjectFilter" yaml:"subjectFilter"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/eventgrid_system_topic_event_subscription#timeouts EventgridSystemTopicEventSubscription#timeouts}
	Timeouts *EventgridSystemTopicEventSubscriptionTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// webhook_endpoint block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/eventgrid_system_topic_event_subscription#webhook_endpoint EventgridSystemTopicEventSubscription#webhook_endpoint}
	WebhookEndpoint *EventgridSystemTopicEventSubscriptionWebhookEndpoint `field:"optional" json:"webhookEndpoint" yaml:"webhookEndpoint"`
}

