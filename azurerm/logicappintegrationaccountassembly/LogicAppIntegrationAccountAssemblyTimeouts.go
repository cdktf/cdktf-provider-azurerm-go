// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package logicappintegrationaccountassembly


type LogicAppIntegrationAccountAssemblyTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.105.0/docs/resources/logic_app_integration_account_assembly#create LogicAppIntegrationAccountAssembly#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.105.0/docs/resources/logic_app_integration_account_assembly#delete LogicAppIntegrationAccountAssembly#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.105.0/docs/resources/logic_app_integration_account_assembly#read LogicAppIntegrationAccountAssembly#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.105.0/docs/resources/logic_app_integration_account_assembly#update LogicAppIntegrationAccountAssembly#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

