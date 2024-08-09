// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package automationpowershell72module


type AutomationPowershell72ModuleModuleLink struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.115.0/docs/resources/automation_powershell72_module#uri AutomationPowershell72Module#uri}.
	Uri *string `field:"required" json:"uri" yaml:"uri"`
	// hash block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.115.0/docs/resources/automation_powershell72_module#hash AutomationPowershell72Module#hash}
	Hash *AutomationPowershell72ModuleModuleLinkHash `field:"optional" json:"hash" yaml:"hash"`
}

