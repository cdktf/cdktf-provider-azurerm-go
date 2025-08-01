// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apimanagementapidiagnostic


type ApiManagementApiDiagnosticFrontendResponseDataMaskingHeaders struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/api_management_api_diagnostic#mode ApiManagementApiDiagnostic#mode}.
	Mode *string `field:"required" json:"mode" yaml:"mode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/api_management_api_diagnostic#value ApiManagementApiDiagnostic#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

