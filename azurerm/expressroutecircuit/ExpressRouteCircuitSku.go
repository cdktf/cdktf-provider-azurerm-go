package expressroutecircuit


type ExpressRouteCircuitSku struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.63.0/docs/resources/express_route_circuit#family ExpressRouteCircuit#family}.
	Family *string `field:"required" json:"family" yaml:"family"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.63.0/docs/resources/express_route_circuit#tier ExpressRouteCircuit#tier}.
	Tier *string `field:"required" json:"tier" yaml:"tier"`
}

