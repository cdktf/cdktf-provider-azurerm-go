package expressroutecircuitpeering


type ExpressRouteCircuitPeeringTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/express_route_circuit_peering#create ExpressRouteCircuitPeering#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/express_route_circuit_peering#delete ExpressRouteCircuitPeering#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/express_route_circuit_peering#read ExpressRouteCircuitPeering#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/express_route_circuit_peering#update ExpressRouteCircuitPeering#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}
