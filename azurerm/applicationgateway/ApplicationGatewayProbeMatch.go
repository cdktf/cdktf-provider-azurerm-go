// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package applicationgateway


type ApplicationGatewayProbeMatch struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.83.0/docs/resources/application_gateway#status_code ApplicationGateway#status_code}.
	StatusCode *[]*string `field:"required" json:"statusCode" yaml:"statusCode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.83.0/docs/resources/application_gateway#body ApplicationGateway#body}.
	Body *string `field:"optional" json:"body" yaml:"body"`
}

