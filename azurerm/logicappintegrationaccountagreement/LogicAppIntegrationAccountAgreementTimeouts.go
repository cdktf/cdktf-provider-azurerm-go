// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package logicappintegrationaccountagreement


type LogicAppIntegrationAccountAgreementTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.85.0/docs/resources/logic_app_integration_account_agreement#create LogicAppIntegrationAccountAgreement#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.85.0/docs/resources/logic_app_integration_account_agreement#delete LogicAppIntegrationAccountAgreement#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.85.0/docs/resources/logic_app_integration_account_agreement#read LogicAppIntegrationAccountAgreement#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.85.0/docs/resources/logic_app_integration_account_agreement#update LogicAppIntegrationAccountAgreement#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

