// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package eventgrideventsubscription


type EventgridEventSubscriptionDeliveryProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/eventgrid_event_subscription#header_name EventgridEventSubscription#header_name}.
	HeaderName *string `field:"required" json:"headerName" yaml:"headerName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/eventgrid_event_subscription#type EventgridEventSubscription#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/eventgrid_event_subscription#secret EventgridEventSubscription#secret}.
	Secret interface{} `field:"optional" json:"secret" yaml:"secret"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/eventgrid_event_subscription#source_field EventgridEventSubscription#source_field}.
	SourceField *string `field:"optional" json:"sourceField" yaml:"sourceField"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/eventgrid_event_subscription#value EventgridEventSubscription#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

