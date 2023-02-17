package cdnfrontdoorrule


type CdnFrontdoorRuleConditionsCookiesCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_frontdoor_rule#cookie_name CdnFrontdoorRule#cookie_name}.
	CookieName *string `field:"required" json:"cookieName" yaml:"cookieName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_frontdoor_rule#operator CdnFrontdoorRule#operator}.
	Operator *string `field:"required" json:"operator" yaml:"operator"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_frontdoor_rule#match_values CdnFrontdoorRule#match_values}.
	MatchValues *[]*string `field:"optional" json:"matchValues" yaml:"matchValues"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_frontdoor_rule#negate_condition CdnFrontdoorRule#negate_condition}.
	NegateCondition interface{} `field:"optional" json:"negateCondition" yaml:"negateCondition"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_frontdoor_rule#transforms CdnFrontdoorRule#transforms}.
	Transforms *[]*string `field:"optional" json:"transforms" yaml:"transforms"`
}

