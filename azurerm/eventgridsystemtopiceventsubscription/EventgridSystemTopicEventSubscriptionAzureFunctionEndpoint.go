// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package eventgridsystemtopiceventsubscription


type EventgridSystemTopicEventSubscriptionAzureFunctionEndpoint struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.16.0/docs/resources/eventgrid_system_topic_event_subscription#function_id EventgridSystemTopicEventSubscription#function_id}.
	FunctionId *string `field:"required" json:"functionId" yaml:"functionId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.16.0/docs/resources/eventgrid_system_topic_event_subscription#max_events_per_batch EventgridSystemTopicEventSubscription#max_events_per_batch}.
	MaxEventsPerBatch *float64 `field:"optional" json:"maxEventsPerBatch" yaml:"maxEventsPerBatch"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.16.0/docs/resources/eventgrid_system_topic_event_subscription#preferred_batch_size_in_kilobytes EventgridSystemTopicEventSubscription#preferred_batch_size_in_kilobytes}.
	PreferredBatchSizeInKilobytes *float64 `field:"optional" json:"preferredBatchSizeInKilobytes" yaml:"preferredBatchSizeInKilobytes"`
}

