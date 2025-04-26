// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package routemap


type RouteMapRuleAction struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.27.0/docs/resources/route_map#type RouteMap#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// parameter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.27.0/docs/resources/route_map#parameter RouteMap#parameter}
	Parameter interface{} `field:"optional" json:"parameter" yaml:"parameter"`
}

