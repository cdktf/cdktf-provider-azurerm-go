package expressrouteport


type ExpressRoutePortIdentity struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/express_route_port#identity_ids ExpressRoutePort#identity_ids}.
	IdentityIds *[]*string `field:"required" json:"identityIds" yaml:"identityIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/express_route_port#type ExpressRoutePort#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}

