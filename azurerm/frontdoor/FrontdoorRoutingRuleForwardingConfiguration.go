// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package frontdoor


type FrontdoorRoutingRuleForwardingConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.84.0/docs/resources/frontdoor#backend_pool_name Frontdoor#backend_pool_name}.
	BackendPoolName *string `field:"required" json:"backendPoolName" yaml:"backendPoolName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.84.0/docs/resources/frontdoor#cache_duration Frontdoor#cache_duration}.
	CacheDuration *string `field:"optional" json:"cacheDuration" yaml:"cacheDuration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.84.0/docs/resources/frontdoor#cache_enabled Frontdoor#cache_enabled}.
	CacheEnabled interface{} `field:"optional" json:"cacheEnabled" yaml:"cacheEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.84.0/docs/resources/frontdoor#cache_query_parameters Frontdoor#cache_query_parameters}.
	CacheQueryParameters *[]*string `field:"optional" json:"cacheQueryParameters" yaml:"cacheQueryParameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.84.0/docs/resources/frontdoor#cache_query_parameter_strip_directive Frontdoor#cache_query_parameter_strip_directive}.
	CacheQueryParameterStripDirective *string `field:"optional" json:"cacheQueryParameterStripDirective" yaml:"cacheQueryParameterStripDirective"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.84.0/docs/resources/frontdoor#cache_use_dynamic_compression Frontdoor#cache_use_dynamic_compression}.
	CacheUseDynamicCompression interface{} `field:"optional" json:"cacheUseDynamicCompression" yaml:"cacheUseDynamicCompression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.84.0/docs/resources/frontdoor#custom_forwarding_path Frontdoor#custom_forwarding_path}.
	CustomForwardingPath *string `field:"optional" json:"customForwardingPath" yaml:"customForwardingPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.84.0/docs/resources/frontdoor#forwarding_protocol Frontdoor#forwarding_protocol}.
	ForwardingProtocol *string `field:"optional" json:"forwardingProtocol" yaml:"forwardingProtocol"`
}

