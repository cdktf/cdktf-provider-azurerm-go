// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package applicationgateway


type ApplicationGatewayRedirectConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/application_gateway#name ApplicationGateway#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/application_gateway#redirect_type ApplicationGateway#redirect_type}.
	RedirectType *string `field:"required" json:"redirectType" yaml:"redirectType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/application_gateway#include_path ApplicationGateway#include_path}.
	IncludePath interface{} `field:"optional" json:"includePath" yaml:"includePath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/application_gateway#include_query_string ApplicationGateway#include_query_string}.
	IncludeQueryString interface{} `field:"optional" json:"includeQueryString" yaml:"includeQueryString"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/application_gateway#target_listener_name ApplicationGateway#target_listener_name}.
	TargetListenerName *string `field:"optional" json:"targetListenerName" yaml:"targetListenerName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/application_gateway#target_url ApplicationGateway#target_url}.
	TargetUrl *string `field:"optional" json:"targetUrl" yaml:"targetUrl"`
}

