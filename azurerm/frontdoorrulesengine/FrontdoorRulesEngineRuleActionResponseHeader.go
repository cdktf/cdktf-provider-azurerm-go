// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package frontdoorrulesengine


type FrontdoorRulesEngineRuleActionResponseHeader struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.33.0/docs/resources/frontdoor_rules_engine#header_action_type FrontdoorRulesEngine#header_action_type}.
	HeaderActionType *string `field:"optional" json:"headerActionType" yaml:"headerActionType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.33.0/docs/resources/frontdoor_rules_engine#header_name FrontdoorRulesEngine#header_name}.
	HeaderName *string `field:"optional" json:"headerName" yaml:"headerName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.33.0/docs/resources/frontdoor_rules_engine#value FrontdoorRulesEngine#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

