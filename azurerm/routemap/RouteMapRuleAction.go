package routemap


type RouteMapRuleAction struct {
	// parameter block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/route_map#parameter RouteMap#parameter}
	Parameter interface{} `field:"required" json:"parameter" yaml:"parameter"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/route_map#type RouteMap#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}

