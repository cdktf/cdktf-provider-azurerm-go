package cdnfrontdoorrule


type CdnFrontdoorRuleConditionsHttpVersionCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_frontdoor_rule#match_values CdnFrontdoorRule#match_values}.
	MatchValues *[]*string `field:"required" json:"matchValues" yaml:"matchValues"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_frontdoor_rule#negate_condition CdnFrontdoorRule#negate_condition}.
	NegateCondition interface{} `field:"optional" json:"negateCondition" yaml:"negateCondition"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_frontdoor_rule#operator CdnFrontdoorRule#operator}.
	Operator *string `field:"optional" json:"operator" yaml:"operator"`
}

