// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apimanagementbackend


type ApiManagementBackendCircuitBreakerRuleFailureConditionStatusCodeRange struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/api_management_backend#max ApiManagementBackend#max}.
	Max *float64 `field:"required" json:"max" yaml:"max"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/api_management_backend#min ApiManagementBackend#min}.
	Min *float64 `field:"required" json:"min" yaml:"min"`
}

