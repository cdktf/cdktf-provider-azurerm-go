// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apimanagementbackend


type ApiManagementBackendCircuitBreakerRuleFailureCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/api_management_backend#interval_duration ApiManagementBackend#interval_duration}.
	IntervalDuration *string `field:"required" json:"intervalDuration" yaml:"intervalDuration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/api_management_backend#count ApiManagementBackend#count}.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/api_management_backend#error_reasons ApiManagementBackend#error_reasons}.
	ErrorReasons *[]*string `field:"optional" json:"errorReasons" yaml:"errorReasons"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/api_management_backend#percentage ApiManagementBackend#percentage}.
	Percentage *float64 `field:"optional" json:"percentage" yaml:"percentage"`
	// status_code_range block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/api_management_backend#status_code_range ApiManagementBackend#status_code_range}
	StatusCodeRange interface{} `field:"optional" json:"statusCodeRange" yaml:"statusCodeRange"`
}

