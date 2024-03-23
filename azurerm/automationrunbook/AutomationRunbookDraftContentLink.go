// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package automationrunbook


type AutomationRunbookDraftContentLink struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.97.1/docs/resources/automation_runbook#uri AutomationRunbook#uri}.
	Uri *string `field:"required" json:"uri" yaml:"uri"`
	// hash block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.97.1/docs/resources/automation_runbook#hash AutomationRunbook#hash}
	Hash *AutomationRunbookDraftContentLinkHash `field:"optional" json:"hash" yaml:"hash"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.97.1/docs/resources/automation_runbook#version AutomationRunbook#version}.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

