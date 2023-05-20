package expressroutecircuitauthorization


type ExpressRouteCircuitAuthorizationTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.57.0/docs/resources/express_route_circuit_authorization#create ExpressRouteCircuitAuthorization#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.57.0/docs/resources/express_route_circuit_authorization#delete ExpressRouteCircuitAuthorization#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.57.0/docs/resources/express_route_circuit_authorization#read ExpressRouteCircuitAuthorization#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.57.0/docs/resources/express_route_circuit_authorization#update ExpressRouteCircuitAuthorization#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

