package expressrouteconnection


type ExpressRouteConnectionRouting struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/express_route_connection#associated_route_table_id ExpressRouteConnection#associated_route_table_id}.
	AssociatedRouteTableId *string `field:"optional" json:"associatedRouteTableId" yaml:"associatedRouteTableId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/express_route_connection#inbound_route_map_id ExpressRouteConnection#inbound_route_map_id}.
	InboundRouteMapId *string `field:"optional" json:"inboundRouteMapId" yaml:"inboundRouteMapId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/express_route_connection#outbound_route_map_id ExpressRouteConnection#outbound_route_map_id}.
	OutboundRouteMapId *string `field:"optional" json:"outboundRouteMapId" yaml:"outboundRouteMapId"`
	// propagated_route_table block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/express_route_connection#propagated_route_table ExpressRouteConnection#propagated_route_table}
	PropagatedRouteTable *ExpressRouteConnectionRoutingPropagatedRouteTable `field:"optional" json:"propagatedRouteTable" yaml:"propagatedRouteTable"`
}

