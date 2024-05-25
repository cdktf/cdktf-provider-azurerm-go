// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cdnfrontdoorrule


type CdnFrontdoorRuleActionsRouteConfigurationOverrideAction struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.105.0/docs/resources/cdn_frontdoor_rule#cache_behavior CdnFrontdoorRule#cache_behavior}.
	CacheBehavior *string `field:"optional" json:"cacheBehavior" yaml:"cacheBehavior"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.105.0/docs/resources/cdn_frontdoor_rule#cache_duration CdnFrontdoorRule#cache_duration}.
	CacheDuration *string `field:"optional" json:"cacheDuration" yaml:"cacheDuration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.105.0/docs/resources/cdn_frontdoor_rule#cdn_frontdoor_origin_group_id CdnFrontdoorRule#cdn_frontdoor_origin_group_id}.
	CdnFrontdoorOriginGroupId *string `field:"optional" json:"cdnFrontdoorOriginGroupId" yaml:"cdnFrontdoorOriginGroupId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.105.0/docs/resources/cdn_frontdoor_rule#compression_enabled CdnFrontdoorRule#compression_enabled}.
	CompressionEnabled interface{} `field:"optional" json:"compressionEnabled" yaml:"compressionEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.105.0/docs/resources/cdn_frontdoor_rule#forwarding_protocol CdnFrontdoorRule#forwarding_protocol}.
	ForwardingProtocol *string `field:"optional" json:"forwardingProtocol" yaml:"forwardingProtocol"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.105.0/docs/resources/cdn_frontdoor_rule#query_string_caching_behavior CdnFrontdoorRule#query_string_caching_behavior}.
	QueryStringCachingBehavior *string `field:"optional" json:"queryStringCachingBehavior" yaml:"queryStringCachingBehavior"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.105.0/docs/resources/cdn_frontdoor_rule#query_string_parameters CdnFrontdoorRule#query_string_parameters}.
	QueryStringParameters *[]*string `field:"optional" json:"queryStringParameters" yaml:"queryStringParameters"`
}

