// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cdnfrontdoorroute

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CdnFrontdoorRouteConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/cdn_frontdoor_route#cdn_frontdoor_endpoint_id CdnFrontdoorRoute#cdn_frontdoor_endpoint_id}.
	CdnFrontdoorEndpointId *string `field:"required" json:"cdnFrontdoorEndpointId" yaml:"cdnFrontdoorEndpointId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/cdn_frontdoor_route#cdn_frontdoor_origin_group_id CdnFrontdoorRoute#cdn_frontdoor_origin_group_id}.
	CdnFrontdoorOriginGroupId *string `field:"required" json:"cdnFrontdoorOriginGroupId" yaml:"cdnFrontdoorOriginGroupId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/cdn_frontdoor_route#cdn_frontdoor_origin_ids CdnFrontdoorRoute#cdn_frontdoor_origin_ids}.
	CdnFrontdoorOriginIds *[]*string `field:"required" json:"cdnFrontdoorOriginIds" yaml:"cdnFrontdoorOriginIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/cdn_frontdoor_route#name CdnFrontdoorRoute#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/cdn_frontdoor_route#patterns_to_match CdnFrontdoorRoute#patterns_to_match}.
	PatternsToMatch *[]*string `field:"required" json:"patternsToMatch" yaml:"patternsToMatch"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/cdn_frontdoor_route#supported_protocols CdnFrontdoorRoute#supported_protocols}.
	SupportedProtocols *[]*string `field:"required" json:"supportedProtocols" yaml:"supportedProtocols"`
	// cache block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/cdn_frontdoor_route#cache CdnFrontdoorRoute#cache}
	Cache *CdnFrontdoorRouteCache `field:"optional" json:"cache" yaml:"cache"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/cdn_frontdoor_route#cdn_frontdoor_custom_domain_ids CdnFrontdoorRoute#cdn_frontdoor_custom_domain_ids}.
	CdnFrontdoorCustomDomainIds *[]*string `field:"optional" json:"cdnFrontdoorCustomDomainIds" yaml:"cdnFrontdoorCustomDomainIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/cdn_frontdoor_route#cdn_frontdoor_origin_path CdnFrontdoorRoute#cdn_frontdoor_origin_path}.
	CdnFrontdoorOriginPath *string `field:"optional" json:"cdnFrontdoorOriginPath" yaml:"cdnFrontdoorOriginPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/cdn_frontdoor_route#cdn_frontdoor_rule_set_ids CdnFrontdoorRoute#cdn_frontdoor_rule_set_ids}.
	CdnFrontdoorRuleSetIds *[]*string `field:"optional" json:"cdnFrontdoorRuleSetIds" yaml:"cdnFrontdoorRuleSetIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/cdn_frontdoor_route#enabled CdnFrontdoorRoute#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/cdn_frontdoor_route#forwarding_protocol CdnFrontdoorRoute#forwarding_protocol}.
	ForwardingProtocol *string `field:"optional" json:"forwardingProtocol" yaml:"forwardingProtocol"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/cdn_frontdoor_route#https_redirect_enabled CdnFrontdoorRoute#https_redirect_enabled}.
	HttpsRedirectEnabled interface{} `field:"optional" json:"httpsRedirectEnabled" yaml:"httpsRedirectEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/cdn_frontdoor_route#id CdnFrontdoorRoute#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/cdn_frontdoor_route#link_to_default_domain CdnFrontdoorRoute#link_to_default_domain}.
	LinkToDefaultDomain interface{} `field:"optional" json:"linkToDefaultDomain" yaml:"linkToDefaultDomain"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/cdn_frontdoor_route#timeouts CdnFrontdoorRoute#timeouts}
	Timeouts *CdnFrontdoorRouteTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

