// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apimanagementapi


type ApiManagementApiOauth2Authorization struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.98.0/docs/resources/api_management_api#authorization_server_name ApiManagementApi#authorization_server_name}.
	AuthorizationServerName *string `field:"required" json:"authorizationServerName" yaml:"authorizationServerName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.98.0/docs/resources/api_management_api#scope ApiManagementApi#scope}.
	Scope *string `field:"optional" json:"scope" yaml:"scope"`
}

