// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package applicationgateway


type ApplicationGatewayCustomErrorConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/application_gateway#custom_error_page_url ApplicationGateway#custom_error_page_url}.
	CustomErrorPageUrl *string `field:"required" json:"customErrorPageUrl" yaml:"customErrorPageUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/application_gateway#status_code ApplicationGateway#status_code}.
	StatusCode *string `field:"required" json:"statusCode" yaml:"statusCode"`
}

