// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package applicationgateway


type ApplicationGatewayRewriteRuleSetRewriteRuleResponseHeaderConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.75.0/docs/resources/application_gateway#header_name ApplicationGateway#header_name}.
	HeaderName *string `field:"required" json:"headerName" yaml:"headerName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.75.0/docs/resources/application_gateway#header_value ApplicationGateway#header_value}.
	HeaderValue *string `field:"required" json:"headerValue" yaml:"headerValue"`
}

