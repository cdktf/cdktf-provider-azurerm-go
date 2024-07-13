// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package eventgrideventsubscription


type EventgridEventSubscriptionAdvancedFilterStringIn struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/eventgrid_event_subscription#key EventgridEventSubscription#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/eventgrid_event_subscription#values EventgridEventSubscription#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

