// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package virtualhubconnection


type VirtualHubConnectionRouting struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.74.0/docs/resources/virtual_hub_connection#associated_route_table_id VirtualHubConnection#associated_route_table_id}.
	AssociatedRouteTableId *string `field:"optional" json:"associatedRouteTableId" yaml:"associatedRouteTableId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.74.0/docs/resources/virtual_hub_connection#inbound_route_map_id VirtualHubConnection#inbound_route_map_id}.
	InboundRouteMapId *string `field:"optional" json:"inboundRouteMapId" yaml:"inboundRouteMapId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.74.0/docs/resources/virtual_hub_connection#outbound_route_map_id VirtualHubConnection#outbound_route_map_id}.
	OutboundRouteMapId *string `field:"optional" json:"outboundRouteMapId" yaml:"outboundRouteMapId"`
	// propagated_route_table block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.74.0/docs/resources/virtual_hub_connection#propagated_route_table VirtualHubConnection#propagated_route_table}
	PropagatedRouteTable *VirtualHubConnectionRoutingPropagatedRouteTable `field:"optional" json:"propagatedRouteTable" yaml:"propagatedRouteTable"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.74.0/docs/resources/virtual_hub_connection#static_vnet_local_route_override_criteria VirtualHubConnection#static_vnet_local_route_override_criteria}.
	StaticVnetLocalRouteOverrideCriteria *string `field:"optional" json:"staticVnetLocalRouteOverrideCriteria" yaml:"staticVnetLocalRouteOverrideCriteria"`
	// static_vnet_route block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.74.0/docs/resources/virtual_hub_connection#static_vnet_route VirtualHubConnection#static_vnet_route}
	StaticVnetRoute interface{} `field:"optional" json:"staticVnetRoute" yaml:"staticVnetRoute"`
}

