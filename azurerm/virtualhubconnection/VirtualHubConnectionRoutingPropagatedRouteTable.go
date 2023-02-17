package virtualhubconnection


type VirtualHubConnectionRoutingPropagatedRouteTable struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_hub_connection#labels VirtualHubConnection#labels}.
	Labels *[]*string `field:"optional" json:"labels" yaml:"labels"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_hub_connection#route_table_ids VirtualHubConnection#route_table_ids}.
	RouteTableIds *[]*string `field:"optional" json:"routeTableIds" yaml:"routeTableIds"`
}

