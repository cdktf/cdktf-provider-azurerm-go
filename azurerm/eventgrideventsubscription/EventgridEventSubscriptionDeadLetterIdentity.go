// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package eventgrideventsubscription


type EventgridEventSubscriptionDeadLetterIdentity struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.3.0/docs/resources/eventgrid_event_subscription#type EventgridEventSubscription#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.3.0/docs/resources/eventgrid_event_subscription#user_assigned_identity EventgridEventSubscription#user_assigned_identity}.
	UserAssignedIdentity *string `field:"optional" json:"userAssignedIdentity" yaml:"userAssignedIdentity"`
}

