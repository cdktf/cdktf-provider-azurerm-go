// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package virtualhubroutingintent


type VirtualHubRoutingIntentRoutingPolicy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.39.0/docs/resources/virtual_hub_routing_intent#destinations VirtualHubRoutingIntent#destinations}.
	Destinations *[]*string `field:"required" json:"destinations" yaml:"destinations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.39.0/docs/resources/virtual_hub_routing_intent#name VirtualHubRoutingIntent#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.39.0/docs/resources/virtual_hub_routing_intent#next_hop VirtualHubRoutingIntent#next_hop}.
	NextHop *string `field:"required" json:"nextHop" yaml:"nextHop"`
}

