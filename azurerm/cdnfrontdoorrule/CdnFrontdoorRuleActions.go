// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cdnfrontdoorrule


type CdnFrontdoorRuleActions struct {
	// request_header_action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.105.0/docs/resources/cdn_frontdoor_rule#request_header_action CdnFrontdoorRule#request_header_action}
	RequestHeaderAction interface{} `field:"optional" json:"requestHeaderAction" yaml:"requestHeaderAction"`
	// response_header_action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.105.0/docs/resources/cdn_frontdoor_rule#response_header_action CdnFrontdoorRule#response_header_action}
	ResponseHeaderAction interface{} `field:"optional" json:"responseHeaderAction" yaml:"responseHeaderAction"`
	// route_configuration_override_action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.105.0/docs/resources/cdn_frontdoor_rule#route_configuration_override_action CdnFrontdoorRule#route_configuration_override_action}
	RouteConfigurationOverrideAction *CdnFrontdoorRuleActionsRouteConfigurationOverrideAction `field:"optional" json:"routeConfigurationOverrideAction" yaml:"routeConfigurationOverrideAction"`
	// url_redirect_action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.105.0/docs/resources/cdn_frontdoor_rule#url_redirect_action CdnFrontdoorRule#url_redirect_action}
	UrlRedirectAction *CdnFrontdoorRuleActionsUrlRedirectAction `field:"optional" json:"urlRedirectAction" yaml:"urlRedirectAction"`
	// url_rewrite_action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.105.0/docs/resources/cdn_frontdoor_rule#url_rewrite_action CdnFrontdoorRule#url_rewrite_action}
	UrlRewriteAction *CdnFrontdoorRuleActionsUrlRewriteAction `field:"optional" json:"urlRewriteAction" yaml:"urlRewriteAction"`
}

