// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apimanagementgatewayapi


type ApiManagementGatewayApiTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/api_management_gateway_api#create ApiManagementGatewayApi#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/api_management_gateway_api#delete ApiManagementGatewayApi#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/api_management_gateway_api#read ApiManagementGatewayApi#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}

