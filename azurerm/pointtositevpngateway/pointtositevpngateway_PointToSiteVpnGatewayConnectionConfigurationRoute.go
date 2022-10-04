package pointtositevpngateway


type PointToSiteVpnGatewayConnectionConfigurationRoute struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/point_to_site_vpn_gateway#associated_route_table_id PointToSiteVpnGateway#associated_route_table_id}.
	AssociatedRouteTableId *string `field:"required" json:"associatedRouteTableId" yaml:"associatedRouteTableId"`
	// propagated_route_table block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/point_to_site_vpn_gateway#propagated_route_table PointToSiteVpnGateway#propagated_route_table}
	PropagatedRouteTable *PointToSiteVpnGatewayConnectionConfigurationRoutePropagatedRouteTable `field:"optional" json:"propagatedRouteTable" yaml:"propagatedRouteTable"`
}

