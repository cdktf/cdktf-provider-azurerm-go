package expressroutecircuit


type ExpressRouteCircuitSku struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/express_route_circuit#family ExpressRouteCircuit#family}.
	Family *string `field:"required" json:"family" yaml:"family"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/express_route_circuit#tier ExpressRouteCircuit#tier}.
	Tier *string `field:"required" json:"tier" yaml:"tier"`
}

