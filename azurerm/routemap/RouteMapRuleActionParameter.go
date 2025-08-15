// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package routemap


type RouteMapRuleActionParameter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.40.0/docs/resources/route_map#as_path RouteMap#as_path}.
	AsPath *[]*string `field:"optional" json:"asPath" yaml:"asPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.40.0/docs/resources/route_map#community RouteMap#community}.
	Community *[]*string `field:"optional" json:"community" yaml:"community"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.40.0/docs/resources/route_map#route_prefix RouteMap#route_prefix}.
	RoutePrefix *[]*string `field:"optional" json:"routePrefix" yaml:"routePrefix"`
}

