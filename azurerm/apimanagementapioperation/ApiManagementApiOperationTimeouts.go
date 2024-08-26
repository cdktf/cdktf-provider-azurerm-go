// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apimanagementapioperation


type ApiManagementApiOperationTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.0.1/docs/resources/api_management_api_operation#create ApiManagementApiOperation#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.0.1/docs/resources/api_management_api_operation#delete ApiManagementApiOperation#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.0.1/docs/resources/api_management_api_operation#read ApiManagementApiOperation#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.0.1/docs/resources/api_management_api_operation#update ApiManagementApiOperation#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

