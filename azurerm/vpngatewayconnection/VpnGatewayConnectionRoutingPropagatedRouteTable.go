package vpngatewayconnection


type VpnGatewayConnectionRoutingPropagatedRouteTable struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/vpn_gateway_connection#route_table_ids VpnGatewayConnection#route_table_ids}.
	RouteTableIds *[]*string `field:"required" json:"routeTableIds" yaml:"routeTableIds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/vpn_gateway_connection#labels VpnGatewayConnection#labels}.
	Labels *[]*string `field:"optional" json:"labels" yaml:"labels"`
}
