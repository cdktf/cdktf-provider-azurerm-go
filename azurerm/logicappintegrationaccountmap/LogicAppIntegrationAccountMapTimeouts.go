// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package logicappintegrationaccountmap


type LogicAppIntegrationAccountMapTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.87.0/docs/resources/logic_app_integration_account_map#create LogicAppIntegrationAccountMap#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.87.0/docs/resources/logic_app_integration_account_map#delete LogicAppIntegrationAccountMap#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.87.0/docs/resources/logic_app_integration_account_map#read LogicAppIntegrationAccountMap#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.87.0/docs/resources/logic_app_integration_account_map#update LogicAppIntegrationAccountMap#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

