// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apimanagementdiagnostic


type ApiManagementDiagnosticBackendResponse struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.72.0/docs/resources/api_management_diagnostic#body_bytes ApiManagementDiagnostic#body_bytes}.
	BodyBytes *float64 `field:"optional" json:"bodyBytes" yaml:"bodyBytes"`
	// data_masking block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.72.0/docs/resources/api_management_diagnostic#data_masking ApiManagementDiagnostic#data_masking}
	DataMasking *ApiManagementDiagnosticBackendResponseDataMasking `field:"optional" json:"dataMasking" yaml:"dataMasking"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.72.0/docs/resources/api_management_diagnostic#headers_to_log ApiManagementDiagnostic#headers_to_log}.
	HeadersToLog *[]*string `field:"optional" json:"headersToLog" yaml:"headersToLog"`
}

