package routemap


type RouteMapRuleAction struct {
	// parameter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.57.0/docs/resources/route_map#parameter RouteMap#parameter}
	Parameter interface{} `field:"required" json:"parameter" yaml:"parameter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.57.0/docs/resources/route_map#type RouteMap#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}

