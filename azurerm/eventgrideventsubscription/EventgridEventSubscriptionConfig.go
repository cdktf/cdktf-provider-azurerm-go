// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package eventgrideventsubscription

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EventgridEventSubscriptionConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/eventgrid_event_subscription#name EventgridEventSubscription#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/eventgrid_event_subscription#scope EventgridEventSubscription#scope}.
	Scope *string `field:"required" json:"scope" yaml:"scope"`
	// advanced_filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/eventgrid_event_subscription#advanced_filter EventgridEventSubscription#advanced_filter}
	AdvancedFilter *EventgridEventSubscriptionAdvancedFilter `field:"optional" json:"advancedFilter" yaml:"advancedFilter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/eventgrid_event_subscription#advanced_filtering_on_arrays_enabled EventgridEventSubscription#advanced_filtering_on_arrays_enabled}.
	AdvancedFilteringOnArraysEnabled interface{} `field:"optional" json:"advancedFilteringOnArraysEnabled" yaml:"advancedFilteringOnArraysEnabled"`
	// azure_function_endpoint block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/eventgrid_event_subscription#azure_function_endpoint EventgridEventSubscription#azure_function_endpoint}
	AzureFunctionEndpoint *EventgridEventSubscriptionAzureFunctionEndpoint `field:"optional" json:"azureFunctionEndpoint" yaml:"azureFunctionEndpoint"`
	// dead_letter_identity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/eventgrid_event_subscription#dead_letter_identity EventgridEventSubscription#dead_letter_identity}
	DeadLetterIdentity *EventgridEventSubscriptionDeadLetterIdentity `field:"optional" json:"deadLetterIdentity" yaml:"deadLetterIdentity"`
	// delivery_identity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/eventgrid_event_subscription#delivery_identity EventgridEventSubscription#delivery_identity}
	DeliveryIdentity *EventgridEventSubscriptionDeliveryIdentity `field:"optional" json:"deliveryIdentity" yaml:"deliveryIdentity"`
	// delivery_property block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/eventgrid_event_subscription#delivery_property EventgridEventSubscription#delivery_property}
	DeliveryProperty interface{} `field:"optional" json:"deliveryProperty" yaml:"deliveryProperty"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/eventgrid_event_subscription#event_delivery_schema EventgridEventSubscription#event_delivery_schema}.
	EventDeliverySchema *string `field:"optional" json:"eventDeliverySchema" yaml:"eventDeliverySchema"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/eventgrid_event_subscription#eventhub_endpoint_id EventgridEventSubscription#eventhub_endpoint_id}.
	EventhubEndpointId *string `field:"optional" json:"eventhubEndpointId" yaml:"eventhubEndpointId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/eventgrid_event_subscription#expiration_time_utc EventgridEventSubscription#expiration_time_utc}.
	ExpirationTimeUtc *string `field:"optional" json:"expirationTimeUtc" yaml:"expirationTimeUtc"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/eventgrid_event_subscription#hybrid_connection_endpoint_id EventgridEventSubscription#hybrid_connection_endpoint_id}.
	HybridConnectionEndpointId *string `field:"optional" json:"hybridConnectionEndpointId" yaml:"hybridConnectionEndpointId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/eventgrid_event_subscription#id EventgridEventSubscription#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/eventgrid_event_subscription#included_event_types EventgridEventSubscription#included_event_types}.
	IncludedEventTypes *[]*string `field:"optional" json:"includedEventTypes" yaml:"includedEventTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/eventgrid_event_subscription#labels EventgridEventSubscription#labels}.
	Labels *[]*string `field:"optional" json:"labels" yaml:"labels"`
	// retry_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/eventgrid_event_subscription#retry_policy EventgridEventSubscription#retry_policy}
	RetryPolicy *EventgridEventSubscriptionRetryPolicy `field:"optional" json:"retryPolicy" yaml:"retryPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/eventgrid_event_subscription#service_bus_queue_endpoint_id EventgridEventSubscription#service_bus_queue_endpoint_id}.
	ServiceBusQueueEndpointId *string `field:"optional" json:"serviceBusQueueEndpointId" yaml:"serviceBusQueueEndpointId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/eventgrid_event_subscription#service_bus_topic_endpoint_id EventgridEventSubscription#service_bus_topic_endpoint_id}.
	ServiceBusTopicEndpointId *string `field:"optional" json:"serviceBusTopicEndpointId" yaml:"serviceBusTopicEndpointId"`
	// storage_blob_dead_letter_destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/eventgrid_event_subscription#storage_blob_dead_letter_destination EventgridEventSubscription#storage_blob_dead_letter_destination}
	StorageBlobDeadLetterDestination *EventgridEventSubscriptionStorageBlobDeadLetterDestination `field:"optional" json:"storageBlobDeadLetterDestination" yaml:"storageBlobDeadLetterDestination"`
	// storage_queue_endpoint block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/eventgrid_event_subscription#storage_queue_endpoint EventgridEventSubscription#storage_queue_endpoint}
	StorageQueueEndpoint *EventgridEventSubscriptionStorageQueueEndpoint `field:"optional" json:"storageQueueEndpoint" yaml:"storageQueueEndpoint"`
	// subject_filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/eventgrid_event_subscription#subject_filter EventgridEventSubscription#subject_filter}
	SubjectFilter *EventgridEventSubscriptionSubjectFilter `field:"optional" json:"subjectFilter" yaml:"subjectFilter"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/eventgrid_event_subscription#timeouts EventgridEventSubscription#timeouts}
	Timeouts *EventgridEventSubscriptionTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// webhook_endpoint block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/eventgrid_event_subscription#webhook_endpoint EventgridEventSubscription#webhook_endpoint}
	WebhookEndpoint *EventgridEventSubscriptionWebhookEndpoint `field:"optional" json:"webhookEndpoint" yaml:"webhookEndpoint"`
}

