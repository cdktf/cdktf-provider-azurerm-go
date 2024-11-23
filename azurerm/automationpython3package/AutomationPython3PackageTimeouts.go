// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package automationpython3package


type AutomationPython3PackageTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/automation_python3_package#create AutomationPython3Package#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/automation_python3_package#delete AutomationPython3Package#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/automation_python3_package#read AutomationPython3Package#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/automation_python3_package#update AutomationPython3Package#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

