// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package automationpowershell72module


type AutomationPowershell72ModuleTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.99.0/docs/resources/automation_powershell72_module#create AutomationPowershell72Module#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.99.0/docs/resources/automation_powershell72_module#delete AutomationPowershell72Module#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.99.0/docs/resources/automation_powershell72_module#read AutomationPowershell72Module#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.99.0/docs/resources/automation_powershell72_module#update AutomationPowershell72Module#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

