// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pointtositevpngateway


type PointToSiteVpnGatewayConnectionConfigurationRoute struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.94.0/docs/resources/point_to_site_vpn_gateway#associated_route_table_id PointToSiteVpnGateway#associated_route_table_id}.
	AssociatedRouteTableId *string `field:"required" json:"associatedRouteTableId" yaml:"associatedRouteTableId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.94.0/docs/resources/point_to_site_vpn_gateway#inbound_route_map_id PointToSiteVpnGateway#inbound_route_map_id}.
	InboundRouteMapId *string `field:"optional" json:"inboundRouteMapId" yaml:"inboundRouteMapId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.94.0/docs/resources/point_to_site_vpn_gateway#outbound_route_map_id PointToSiteVpnGateway#outbound_route_map_id}.
	OutboundRouteMapId *string `field:"optional" json:"outboundRouteMapId" yaml:"outboundRouteMapId"`
	// propagated_route_table block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.94.0/docs/resources/point_to_site_vpn_gateway#propagated_route_table PointToSiteVpnGateway#propagated_route_table}
	PropagatedRouteTable *PointToSiteVpnGatewayConnectionConfigurationRoutePropagatedRouteTable `field:"optional" json:"propagatedRouteTable" yaml:"propagatedRouteTable"`
}

