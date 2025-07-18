// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package frontdoor


type FrontdoorRoutingRuleRedirectConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.37.0/docs/resources/frontdoor#redirect_protocol Frontdoor#redirect_protocol}.
	RedirectProtocol *string `field:"required" json:"redirectProtocol" yaml:"redirectProtocol"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.37.0/docs/resources/frontdoor#redirect_type Frontdoor#redirect_type}.
	RedirectType *string `field:"required" json:"redirectType" yaml:"redirectType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.37.0/docs/resources/frontdoor#custom_fragment Frontdoor#custom_fragment}.
	CustomFragment *string `field:"optional" json:"customFragment" yaml:"customFragment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.37.0/docs/resources/frontdoor#custom_host Frontdoor#custom_host}.
	CustomHost *string `field:"optional" json:"customHost" yaml:"customHost"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.37.0/docs/resources/frontdoor#custom_path Frontdoor#custom_path}.
	CustomPath *string `field:"optional" json:"customPath" yaml:"customPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.37.0/docs/resources/frontdoor#custom_query_string Frontdoor#custom_query_string}.
	CustomQueryString *string `field:"optional" json:"customQueryString" yaml:"customQueryString"`
}

