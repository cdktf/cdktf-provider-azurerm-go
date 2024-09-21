// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apimanagementapioperation


type ApiManagementApiOperationResponse struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.3.0/docs/resources/api_management_api_operation#status_code ApiManagementApiOperation#status_code}.
	StatusCode *float64 `field:"required" json:"statusCode" yaml:"statusCode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.3.0/docs/resources/api_management_api_operation#description ApiManagementApiOperation#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// header block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.3.0/docs/resources/api_management_api_operation#header ApiManagementApiOperation#header}
	Header interface{} `field:"optional" json:"header" yaml:"header"`
	// representation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.3.0/docs/resources/api_management_api_operation#representation ApiManagementApiOperation#representation}
	Representation interface{} `field:"optional" json:"representation" yaml:"representation"`
}

