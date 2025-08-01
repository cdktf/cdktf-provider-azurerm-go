// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apimanagementbackend


type ApiManagementBackendCredentials struct {
	// authorization block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/api_management_backend#authorization ApiManagementBackend#authorization}
	Authorization *ApiManagementBackendCredentialsAuthorization `field:"optional" json:"authorization" yaml:"authorization"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/api_management_backend#certificate ApiManagementBackend#certificate}.
	Certificate *[]*string `field:"optional" json:"certificate" yaml:"certificate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/api_management_backend#header ApiManagementBackend#header}.
	Header *map[string]*string `field:"optional" json:"header" yaml:"header"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/api_management_backend#query ApiManagementBackend#query}.
	Query *map[string]*string `field:"optional" json:"query" yaml:"query"`
}

