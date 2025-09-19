// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apimanagementdiagnostic


type ApiManagementDiagnosticTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/api_management_diagnostic#create ApiManagementDiagnostic#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/api_management_diagnostic#delete ApiManagementDiagnostic#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/api_management_diagnostic#read ApiManagementDiagnostic#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/api_management_diagnostic#update ApiManagementDiagnostic#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

