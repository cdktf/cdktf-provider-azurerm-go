package routemap


type RouteMapRule struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/route_map#name RouteMap#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// action block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/route_map#action RouteMap#action}
	Action interface{} `field:"optional" json:"action" yaml:"action"`
	// match_criterion block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/route_map#match_criterion RouteMap#match_criterion}
	MatchCriterion interface{} `field:"optional" json:"matchCriterion" yaml:"matchCriterion"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/route_map#next_step_if_matched RouteMap#next_step_if_matched}.
	NextStepIfMatched *string `field:"optional" json:"nextStepIfMatched" yaml:"nextStepIfMatched"`
}

