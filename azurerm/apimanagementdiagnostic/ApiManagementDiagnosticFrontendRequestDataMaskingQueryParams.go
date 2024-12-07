// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apimanagementdiagnostic


type ApiManagementDiagnosticFrontendRequestDataMaskingQueryParams struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/api_management_diagnostic#mode ApiManagementDiagnostic#mode}.
	Mode *string `field:"required" json:"mode" yaml:"mode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/api_management_diagnostic#value ApiManagementDiagnostic#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

