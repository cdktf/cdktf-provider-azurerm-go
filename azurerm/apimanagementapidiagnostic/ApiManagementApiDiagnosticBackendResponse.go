// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apimanagementapidiagnostic


type ApiManagementApiDiagnosticBackendResponse struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.39.0/docs/resources/api_management_api_diagnostic#body_bytes ApiManagementApiDiagnostic#body_bytes}.
	BodyBytes *float64 `field:"optional" json:"bodyBytes" yaml:"bodyBytes"`
	// data_masking block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.39.0/docs/resources/api_management_api_diagnostic#data_masking ApiManagementApiDiagnostic#data_masking}
	DataMasking *ApiManagementApiDiagnosticBackendResponseDataMasking `field:"optional" json:"dataMasking" yaml:"dataMasking"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.39.0/docs/resources/api_management_api_diagnostic#headers_to_log ApiManagementApiDiagnostic#headers_to_log}.
	HeadersToLog *[]*string `field:"optional" json:"headersToLog" yaml:"headersToLog"`
}

