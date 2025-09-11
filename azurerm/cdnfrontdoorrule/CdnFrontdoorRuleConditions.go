// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cdnfrontdoorrule


type CdnFrontdoorRuleConditions struct {
	// client_port_condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs/resources/cdn_frontdoor_rule#client_port_condition CdnFrontdoorRule#client_port_condition}
	ClientPortCondition interface{} `field:"optional" json:"clientPortCondition" yaml:"clientPortCondition"`
	// cookies_condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs/resources/cdn_frontdoor_rule#cookies_condition CdnFrontdoorRule#cookies_condition}
	CookiesCondition interface{} `field:"optional" json:"cookiesCondition" yaml:"cookiesCondition"`
	// host_name_condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs/resources/cdn_frontdoor_rule#host_name_condition CdnFrontdoorRule#host_name_condition}
	HostNameCondition interface{} `field:"optional" json:"hostNameCondition" yaml:"hostNameCondition"`
	// http_version_condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs/resources/cdn_frontdoor_rule#http_version_condition CdnFrontdoorRule#http_version_condition}
	HttpVersionCondition interface{} `field:"optional" json:"httpVersionCondition" yaml:"httpVersionCondition"`
	// is_device_condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs/resources/cdn_frontdoor_rule#is_device_condition CdnFrontdoorRule#is_device_condition}
	IsDeviceCondition interface{} `field:"optional" json:"isDeviceCondition" yaml:"isDeviceCondition"`
	// post_args_condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs/resources/cdn_frontdoor_rule#post_args_condition CdnFrontdoorRule#post_args_condition}
	PostArgsCondition interface{} `field:"optional" json:"postArgsCondition" yaml:"postArgsCondition"`
	// query_string_condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs/resources/cdn_frontdoor_rule#query_string_condition CdnFrontdoorRule#query_string_condition}
	QueryStringCondition interface{} `field:"optional" json:"queryStringCondition" yaml:"queryStringCondition"`
	// remote_address_condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs/resources/cdn_frontdoor_rule#remote_address_condition CdnFrontdoorRule#remote_address_condition}
	RemoteAddressCondition interface{} `field:"optional" json:"remoteAddressCondition" yaml:"remoteAddressCondition"`
	// request_body_condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs/resources/cdn_frontdoor_rule#request_body_condition CdnFrontdoorRule#request_body_condition}
	RequestBodyCondition interface{} `field:"optional" json:"requestBodyCondition" yaml:"requestBodyCondition"`
	// request_header_condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs/resources/cdn_frontdoor_rule#request_header_condition CdnFrontdoorRule#request_header_condition}
	RequestHeaderCondition interface{} `field:"optional" json:"requestHeaderCondition" yaml:"requestHeaderCondition"`
	// request_method_condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs/resources/cdn_frontdoor_rule#request_method_condition CdnFrontdoorRule#request_method_condition}
	RequestMethodCondition interface{} `field:"optional" json:"requestMethodCondition" yaml:"requestMethodCondition"`
	// request_scheme_condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs/resources/cdn_frontdoor_rule#request_scheme_condition CdnFrontdoorRule#request_scheme_condition}
	RequestSchemeCondition interface{} `field:"optional" json:"requestSchemeCondition" yaml:"requestSchemeCondition"`
	// request_uri_condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs/resources/cdn_frontdoor_rule#request_uri_condition CdnFrontdoorRule#request_uri_condition}
	RequestUriCondition interface{} `field:"optional" json:"requestUriCondition" yaml:"requestUriCondition"`
	// server_port_condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs/resources/cdn_frontdoor_rule#server_port_condition CdnFrontdoorRule#server_port_condition}
	ServerPortCondition interface{} `field:"optional" json:"serverPortCondition" yaml:"serverPortCondition"`
	// socket_address_condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs/resources/cdn_frontdoor_rule#socket_address_condition CdnFrontdoorRule#socket_address_condition}
	SocketAddressCondition interface{} `field:"optional" json:"socketAddressCondition" yaml:"socketAddressCondition"`
	// ssl_protocol_condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs/resources/cdn_frontdoor_rule#ssl_protocol_condition CdnFrontdoorRule#ssl_protocol_condition}
	SslProtocolCondition interface{} `field:"optional" json:"sslProtocolCondition" yaml:"sslProtocolCondition"`
	// url_file_extension_condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs/resources/cdn_frontdoor_rule#url_file_extension_condition CdnFrontdoorRule#url_file_extension_condition}
	UrlFileExtensionCondition interface{} `field:"optional" json:"urlFileExtensionCondition" yaml:"urlFileExtensionCondition"`
	// url_filename_condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs/resources/cdn_frontdoor_rule#url_filename_condition CdnFrontdoorRule#url_filename_condition}
	UrlFilenameCondition interface{} `field:"optional" json:"urlFilenameCondition" yaml:"urlFilenameCondition"`
	// url_path_condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs/resources/cdn_frontdoor_rule#url_path_condition CdnFrontdoorRule#url_path_condition}
	UrlPathCondition interface{} `field:"optional" json:"urlPathCondition" yaml:"urlPathCondition"`
}

