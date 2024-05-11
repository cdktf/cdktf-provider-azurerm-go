// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package automationrunbook


type AutomationRunbookDraftParameters struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/automation_runbook#key AutomationRunbook#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/automation_runbook#type AutomationRunbook#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/automation_runbook#default_value AutomationRunbook#default_value}.
	DefaultValue *string `field:"optional" json:"defaultValue" yaml:"defaultValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/automation_runbook#mandatory AutomationRunbook#mandatory}.
	Mandatory interface{} `field:"optional" json:"mandatory" yaml:"mandatory"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/automation_runbook#position AutomationRunbook#position}.
	Position *float64 `field:"optional" json:"position" yaml:"position"`
}

