package routefilter


type RouteFilterRule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.63.0/docs/resources/route_filter#access RouteFilter#access}.
	Access *string `field:"optional" json:"access" yaml:"access"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.63.0/docs/resources/route_filter#communities RouteFilter#communities}.
	Communities *[]*string `field:"optional" json:"communities" yaml:"communities"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.63.0/docs/resources/route_filter#name RouteFilter#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.63.0/docs/resources/route_filter#rule_type RouteFilter#rule_type}.
	RuleType *string `field:"optional" json:"ruleType" yaml:"ruleType"`
}

