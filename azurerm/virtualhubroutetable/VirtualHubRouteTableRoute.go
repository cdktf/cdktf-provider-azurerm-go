// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package virtualhubroutetable


type VirtualHubRouteTableRoute struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.30.0/docs/resources/virtual_hub_route_table#destinations VirtualHubRouteTable#destinations}.
	Destinations *[]*string `field:"required" json:"destinations" yaml:"destinations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.30.0/docs/resources/virtual_hub_route_table#destinations_type VirtualHubRouteTable#destinations_type}.
	DestinationsType *string `field:"required" json:"destinationsType" yaml:"destinationsType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.30.0/docs/resources/virtual_hub_route_table#name VirtualHubRouteTable#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.30.0/docs/resources/virtual_hub_route_table#next_hop VirtualHubRouteTable#next_hop}.
	NextHop *string `field:"required" json:"nextHop" yaml:"nextHop"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.30.0/docs/resources/virtual_hub_route_table#next_hop_type VirtualHubRouteTable#next_hop_type}.
	NextHopType *string `field:"optional" json:"nextHopType" yaml:"nextHopType"`
}

