package expressrouteconnection


type ExpressRouteConnectionRouting struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/express_route_connection#associated_route_table_id ExpressRouteConnection#associated_route_table_id}.
	AssociatedRouteTableId *string `field:"optional" json:"associatedRouteTableId" yaml:"associatedRouteTableId"`
	// propagated_route_table block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/express_route_connection#propagated_route_table ExpressRouteConnection#propagated_route_table}
	PropagatedRouteTable *ExpressRouteConnectionRoutingPropagatedRouteTable `field:"optional" json:"propagatedRouteTable" yaml:"propagatedRouteTable"`
}

