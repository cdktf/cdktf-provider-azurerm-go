// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package securitycenterautomation


type SecurityCenterAutomationSourceRuleSetRule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.107.0/docs/resources/security_center_automation#expected_value SecurityCenterAutomation#expected_value}.
	ExpectedValue *string `field:"required" json:"expectedValue" yaml:"expectedValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.107.0/docs/resources/security_center_automation#operator SecurityCenterAutomation#operator}.
	Operator *string `field:"required" json:"operator" yaml:"operator"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.107.0/docs/resources/security_center_automation#property_path SecurityCenterAutomation#property_path}.
	PropertyPath *string `field:"required" json:"propertyPath" yaml:"propertyPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.107.0/docs/resources/security_center_automation#property_type SecurityCenterAutomation#property_type}.
	PropertyType *string `field:"required" json:"propertyType" yaml:"propertyType"`
}

