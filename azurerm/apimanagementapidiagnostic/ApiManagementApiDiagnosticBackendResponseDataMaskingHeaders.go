// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apimanagementapidiagnostic


type ApiManagementApiDiagnosticBackendResponseDataMaskingHeaders struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.101.0/docs/resources/api_management_api_diagnostic#mode ApiManagementApiDiagnostic#mode}.
	Mode *string `field:"required" json:"mode" yaml:"mode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.101.0/docs/resources/api_management_api_diagnostic#value ApiManagementApiDiagnostic#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

