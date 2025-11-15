// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apimanagementgateway


type ApiManagementGatewayLocationData struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.53.0/docs/resources/api_management_gateway#name ApiManagementGateway#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.53.0/docs/resources/api_management_gateway#city ApiManagementGateway#city}.
	City *string `field:"optional" json:"city" yaml:"city"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.53.0/docs/resources/api_management_gateway#district ApiManagementGateway#district}.
	District *string `field:"optional" json:"district" yaml:"district"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.53.0/docs/resources/api_management_gateway#region ApiManagementGateway#region}.
	Region *string `field:"optional" json:"region" yaml:"region"`
}

