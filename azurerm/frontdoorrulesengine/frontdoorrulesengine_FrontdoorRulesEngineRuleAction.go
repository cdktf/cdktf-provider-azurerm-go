package frontdoorrulesengine


type FrontdoorRulesEngineRuleAction struct {
	// request_header block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/frontdoor_rules_engine#request_header FrontdoorRulesEngine#request_header}
	RequestHeader interface{} `field:"optional" json:"requestHeader" yaml:"requestHeader"`
	// response_header block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/frontdoor_rules_engine#response_header FrontdoorRulesEngine#response_header}
	ResponseHeader interface{} `field:"optional" json:"responseHeader" yaml:"responseHeader"`
}
