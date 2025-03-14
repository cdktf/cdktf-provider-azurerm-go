// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package routemap


type RouteMapRule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.23.0/docs/resources/route_map#name RouteMap#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.23.0/docs/resources/route_map#action RouteMap#action}
	Action interface{} `field:"optional" json:"action" yaml:"action"`
	// match_criterion block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.23.0/docs/resources/route_map#match_criterion RouteMap#match_criterion}
	MatchCriterion interface{} `field:"optional" json:"matchCriterion" yaml:"matchCriterion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.23.0/docs/resources/route_map#next_step_if_matched RouteMap#next_step_if_matched}.
	NextStepIfMatched *string `field:"optional" json:"nextStepIfMatched" yaml:"nextStepIfMatched"`
}

