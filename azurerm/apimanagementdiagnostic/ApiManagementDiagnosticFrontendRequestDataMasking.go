// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apimanagementdiagnostic


type ApiManagementDiagnosticFrontendRequestDataMasking struct {
	// headers block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/resources/api_management_diagnostic#headers ApiManagementDiagnostic#headers}
	Headers interface{} `field:"optional" json:"headers" yaml:"headers"`
	// query_params block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/resources/api_management_diagnostic#query_params ApiManagementDiagnostic#query_params}
	QueryParams interface{} `field:"optional" json:"queryParams" yaml:"queryParams"`
}

