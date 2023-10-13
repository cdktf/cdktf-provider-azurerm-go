// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package eventgridsystemtopiceventsubscription


type EventgridSystemTopicEventSubscriptionStorageQueueEndpoint struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/resources/eventgrid_system_topic_event_subscription#queue_name EventgridSystemTopicEventSubscription#queue_name}.
	QueueName *string `field:"required" json:"queueName" yaml:"queueName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/resources/eventgrid_system_topic_event_subscription#storage_account_id EventgridSystemTopicEventSubscription#storage_account_id}.
	StorageAccountId *string `field:"required" json:"storageAccountId" yaml:"storageAccountId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/resources/eventgrid_system_topic_event_subscription#queue_message_time_to_live_in_seconds EventgridSystemTopicEventSubscription#queue_message_time_to_live_in_seconds}.
	QueueMessageTimeToLiveInSeconds *float64 `field:"optional" json:"queueMessageTimeToLiveInSeconds" yaml:"queueMessageTimeToLiveInSeconds"`
}

