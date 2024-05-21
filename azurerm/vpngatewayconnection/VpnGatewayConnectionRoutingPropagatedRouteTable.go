// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vpngatewayconnection


type VpnGatewayConnectionRoutingPropagatedRouteTable struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.2/docs/resources/vpn_gateway_connection#route_table_ids VpnGatewayConnection#route_table_ids}.
	RouteTableIds *[]*string `field:"required" json:"routeTableIds" yaml:"routeTableIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.2/docs/resources/vpn_gateway_connection#labels VpnGatewayConnection#labels}.
	Labels *[]*string `field:"optional" json:"labels" yaml:"labels"`
}

