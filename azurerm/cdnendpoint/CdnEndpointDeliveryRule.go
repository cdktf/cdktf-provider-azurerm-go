package cdnendpoint


type CdnEndpointDeliveryRule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.56.0/docs/resources/cdn_endpoint#name CdnEndpoint#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.56.0/docs/resources/cdn_endpoint#order CdnEndpoint#order}.
	Order *float64 `field:"required" json:"order" yaml:"order"`
	// cache_expiration_action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.56.0/docs/resources/cdn_endpoint#cache_expiration_action CdnEndpoint#cache_expiration_action}
	CacheExpirationAction *CdnEndpointDeliveryRuleCacheExpirationAction `field:"optional" json:"cacheExpirationAction" yaml:"cacheExpirationAction"`
	// cache_key_query_string_action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.56.0/docs/resources/cdn_endpoint#cache_key_query_string_action CdnEndpoint#cache_key_query_string_action}
	CacheKeyQueryStringAction *CdnEndpointDeliveryRuleCacheKeyQueryStringAction `field:"optional" json:"cacheKeyQueryStringAction" yaml:"cacheKeyQueryStringAction"`
	// cookies_condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.56.0/docs/resources/cdn_endpoint#cookies_condition CdnEndpoint#cookies_condition}
	CookiesCondition interface{} `field:"optional" json:"cookiesCondition" yaml:"cookiesCondition"`
	// device_condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.56.0/docs/resources/cdn_endpoint#device_condition CdnEndpoint#device_condition}
	DeviceCondition *CdnEndpointDeliveryRuleDeviceCondition `field:"optional" json:"deviceCondition" yaml:"deviceCondition"`
	// http_version_condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.56.0/docs/resources/cdn_endpoint#http_version_condition CdnEndpoint#http_version_condition}
	HttpVersionCondition interface{} `field:"optional" json:"httpVersionCondition" yaml:"httpVersionCondition"`
	// modify_request_header_action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.56.0/docs/resources/cdn_endpoint#modify_request_header_action CdnEndpoint#modify_request_header_action}
	ModifyRequestHeaderAction interface{} `field:"optional" json:"modifyRequestHeaderAction" yaml:"modifyRequestHeaderAction"`
	// modify_response_header_action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.56.0/docs/resources/cdn_endpoint#modify_response_header_action CdnEndpoint#modify_response_header_action}
	ModifyResponseHeaderAction interface{} `field:"optional" json:"modifyResponseHeaderAction" yaml:"modifyResponseHeaderAction"`
	// post_arg_condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.56.0/docs/resources/cdn_endpoint#post_arg_condition CdnEndpoint#post_arg_condition}
	PostArgCondition interface{} `field:"optional" json:"postArgCondition" yaml:"postArgCondition"`
	// query_string_condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.56.0/docs/resources/cdn_endpoint#query_string_condition CdnEndpoint#query_string_condition}
	QueryStringCondition interface{} `field:"optional" json:"queryStringCondition" yaml:"queryStringCondition"`
	// remote_address_condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.56.0/docs/resources/cdn_endpoint#remote_address_condition CdnEndpoint#remote_address_condition}
	RemoteAddressCondition interface{} `field:"optional" json:"remoteAddressCondition" yaml:"remoteAddressCondition"`
	// request_body_condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.56.0/docs/resources/cdn_endpoint#request_body_condition CdnEndpoint#request_body_condition}
	RequestBodyCondition interface{} `field:"optional" json:"requestBodyCondition" yaml:"requestBodyCondition"`
	// request_header_condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.56.0/docs/resources/cdn_endpoint#request_header_condition CdnEndpoint#request_header_condition}
	RequestHeaderCondition interface{} `field:"optional" json:"requestHeaderCondition" yaml:"requestHeaderCondition"`
	// request_method_condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.56.0/docs/resources/cdn_endpoint#request_method_condition CdnEndpoint#request_method_condition}
	RequestMethodCondition *CdnEndpointDeliveryRuleRequestMethodCondition `field:"optional" json:"requestMethodCondition" yaml:"requestMethodCondition"`
	// request_scheme_condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.56.0/docs/resources/cdn_endpoint#request_scheme_condition CdnEndpoint#request_scheme_condition}
	RequestSchemeCondition *CdnEndpointDeliveryRuleRequestSchemeCondition `field:"optional" json:"requestSchemeCondition" yaml:"requestSchemeCondition"`
	// request_uri_condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.56.0/docs/resources/cdn_endpoint#request_uri_condition CdnEndpoint#request_uri_condition}
	RequestUriCondition interface{} `field:"optional" json:"requestUriCondition" yaml:"requestUriCondition"`
	// url_file_extension_condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.56.0/docs/resources/cdn_endpoint#url_file_extension_condition CdnEndpoint#url_file_extension_condition}
	UrlFileExtensionCondition interface{} `field:"optional" json:"urlFileExtensionCondition" yaml:"urlFileExtensionCondition"`
	// url_file_name_condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.56.0/docs/resources/cdn_endpoint#url_file_name_condition CdnEndpoint#url_file_name_condition}
	UrlFileNameCondition interface{} `field:"optional" json:"urlFileNameCondition" yaml:"urlFileNameCondition"`
	// url_path_condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.56.0/docs/resources/cdn_endpoint#url_path_condition CdnEndpoint#url_path_condition}
	UrlPathCondition interface{} `field:"optional" json:"urlPathCondition" yaml:"urlPathCondition"`
	// url_redirect_action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.56.0/docs/resources/cdn_endpoint#url_redirect_action CdnEndpoint#url_redirect_action}
	UrlRedirectAction *CdnEndpointDeliveryRuleUrlRedirectAction `field:"optional" json:"urlRedirectAction" yaml:"urlRedirectAction"`
	// url_rewrite_action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.56.0/docs/resources/cdn_endpoint#url_rewrite_action CdnEndpoint#url_rewrite_action}
	UrlRewriteAction *CdnEndpointDeliveryRuleUrlRewriteAction `field:"optional" json:"urlRewriteAction" yaml:"urlRewriteAction"`
}

