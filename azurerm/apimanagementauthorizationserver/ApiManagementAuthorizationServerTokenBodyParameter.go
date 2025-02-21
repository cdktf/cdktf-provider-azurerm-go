// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apimanagementauthorizationserver


type ApiManagementAuthorizationServerTokenBodyParameter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.20.0/docs/resources/api_management_authorization_server#name ApiManagementAuthorizationServer#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.20.0/docs/resources/api_management_authorization_server#value ApiManagementAuthorizationServer#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

