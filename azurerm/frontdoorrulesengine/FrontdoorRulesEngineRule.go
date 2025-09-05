// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package frontdoorrulesengine


type FrontdoorRulesEngineRule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/frontdoor_rules_engine#name FrontdoorRulesEngine#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/frontdoor_rules_engine#priority FrontdoorRulesEngine#priority}.
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
	// action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/frontdoor_rules_engine#action FrontdoorRulesEngine#action}
	Action *FrontdoorRulesEngineRuleAction `field:"optional" json:"action" yaml:"action"`
	// match_condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/frontdoor_rules_engine#match_condition FrontdoorRulesEngine#match_condition}
	MatchCondition interface{} `field:"optional" json:"matchCondition" yaml:"matchCondition"`
}

