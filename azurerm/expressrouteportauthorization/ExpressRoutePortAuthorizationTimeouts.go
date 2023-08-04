package expressrouteportauthorization


type ExpressRoutePortAuthorizationTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.68.0/docs/resources/express_route_port_authorization#create ExpressRoutePortAuthorization#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.68.0/docs/resources/express_route_port_authorization#delete ExpressRoutePortAuthorization#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.68.0/docs/resources/express_route_port_authorization#read ExpressRoutePortAuthorization#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.68.0/docs/resources/express_route_port_authorization#update ExpressRoutePortAuthorization#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

