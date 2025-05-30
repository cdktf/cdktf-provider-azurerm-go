// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apimanagementapioperation


type ApiManagementApiOperationRequest struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.31.0/docs/resources/api_management_api_operation#description ApiManagementApiOperation#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// header block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.31.0/docs/resources/api_management_api_operation#header ApiManagementApiOperation#header}
	Header interface{} `field:"optional" json:"header" yaml:"header"`
	// query_parameter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.31.0/docs/resources/api_management_api_operation#query_parameter ApiManagementApiOperation#query_parameter}
	QueryParameter interface{} `field:"optional" json:"queryParameter" yaml:"queryParameter"`
	// representation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.31.0/docs/resources/api_management_api_operation#representation ApiManagementApiOperation#representation}
	Representation interface{} `field:"optional" json:"representation" yaml:"representation"`
}

