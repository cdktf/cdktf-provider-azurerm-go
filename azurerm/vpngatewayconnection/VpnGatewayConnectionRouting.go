// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vpngatewayconnection


type VpnGatewayConnectionRouting struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.23.0/docs/resources/vpn_gateway_connection#associated_route_table VpnGatewayConnection#associated_route_table}.
	AssociatedRouteTable *string `field:"required" json:"associatedRouteTable" yaml:"associatedRouteTable"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.23.0/docs/resources/vpn_gateway_connection#inbound_route_map_id VpnGatewayConnection#inbound_route_map_id}.
	InboundRouteMapId *string `field:"optional" json:"inboundRouteMapId" yaml:"inboundRouteMapId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.23.0/docs/resources/vpn_gateway_connection#outbound_route_map_id VpnGatewayConnection#outbound_route_map_id}.
	OutboundRouteMapId *string `field:"optional" json:"outboundRouteMapId" yaml:"outboundRouteMapId"`
	// propagated_route_table block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.23.0/docs/resources/vpn_gateway_connection#propagated_route_table VpnGatewayConnection#propagated_route_table}
	PropagatedRouteTable *VpnGatewayConnectionRoutingPropagatedRouteTable `field:"optional" json:"propagatedRouteTable" yaml:"propagatedRouteTable"`
}

