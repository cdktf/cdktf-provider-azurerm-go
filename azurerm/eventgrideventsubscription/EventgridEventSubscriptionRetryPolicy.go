// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package eventgrideventsubscription


type EventgridEventSubscriptionRetryPolicy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/eventgrid_event_subscription#event_time_to_live EventgridEventSubscription#event_time_to_live}.
	EventTimeToLive *float64 `field:"required" json:"eventTimeToLive" yaml:"eventTimeToLive"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/eventgrid_event_subscription#max_delivery_attempts EventgridEventSubscription#max_delivery_attempts}.
	MaxDeliveryAttempts *float64 `field:"required" json:"maxDeliveryAttempts" yaml:"maxDeliveryAttempts"`
}

