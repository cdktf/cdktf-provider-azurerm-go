// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cdnendpoint


type CdnEndpointGlobalDeliveryRule struct {
	// cache_expiration_action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.14.0/docs/resources/cdn_endpoint#cache_expiration_action CdnEndpoint#cache_expiration_action}
	CacheExpirationAction *CdnEndpointGlobalDeliveryRuleCacheExpirationAction `field:"optional" json:"cacheExpirationAction" yaml:"cacheExpirationAction"`
	// cache_key_query_string_action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.14.0/docs/resources/cdn_endpoint#cache_key_query_string_action CdnEndpoint#cache_key_query_string_action}
	CacheKeyQueryStringAction *CdnEndpointGlobalDeliveryRuleCacheKeyQueryStringAction `field:"optional" json:"cacheKeyQueryStringAction" yaml:"cacheKeyQueryStringAction"`
	// modify_request_header_action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.14.0/docs/resources/cdn_endpoint#modify_request_header_action CdnEndpoint#modify_request_header_action}
	ModifyRequestHeaderAction interface{} `field:"optional" json:"modifyRequestHeaderAction" yaml:"modifyRequestHeaderAction"`
	// modify_response_header_action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.14.0/docs/resources/cdn_endpoint#modify_response_header_action CdnEndpoint#modify_response_header_action}
	ModifyResponseHeaderAction interface{} `field:"optional" json:"modifyResponseHeaderAction" yaml:"modifyResponseHeaderAction"`
	// url_redirect_action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.14.0/docs/resources/cdn_endpoint#url_redirect_action CdnEndpoint#url_redirect_action}
	UrlRedirectAction *CdnEndpointGlobalDeliveryRuleUrlRedirectAction `field:"optional" json:"urlRedirectAction" yaml:"urlRedirectAction"`
	// url_rewrite_action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.14.0/docs/resources/cdn_endpoint#url_rewrite_action CdnEndpoint#url_rewrite_action}
	UrlRewriteAction *CdnEndpointGlobalDeliveryRuleUrlRewriteAction `field:"optional" json:"urlRewriteAction" yaml:"urlRewriteAction"`
}

