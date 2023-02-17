package expressrouteconnection


type ExpressRouteConnectionRoutingPropagatedRouteTable struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/express_route_connection#labels ExpressRouteConnection#labels}.
	Labels *[]*string `field:"optional" json:"labels" yaml:"labels"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/express_route_connection#route_table_ids ExpressRouteConnection#route_table_ids}.
	RouteTableIds *[]*string `field:"optional" json:"routeTableIds" yaml:"routeTableIds"`
}

