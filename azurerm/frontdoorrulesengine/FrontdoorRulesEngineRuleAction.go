// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package frontdoorrulesengine


type FrontdoorRulesEngineRuleAction struct {
	// request_header block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/frontdoor_rules_engine#request_header FrontdoorRulesEngine#request_header}
	RequestHeader interface{} `field:"optional" json:"requestHeader" yaml:"requestHeader"`
	// response_header block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/frontdoor_rules_engine#response_header FrontdoorRulesEngine#response_header}
	ResponseHeader interface{} `field:"optional" json:"responseHeader" yaml:"responseHeader"`
}

