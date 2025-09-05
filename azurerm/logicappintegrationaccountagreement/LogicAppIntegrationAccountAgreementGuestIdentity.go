// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package logicappintegrationaccountagreement


type LogicAppIntegrationAccountAgreementGuestIdentity struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/logic_app_integration_account_agreement#qualifier LogicAppIntegrationAccountAgreement#qualifier}.
	Qualifier *string `field:"required" json:"qualifier" yaml:"qualifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/logic_app_integration_account_agreement#value LogicAppIntegrationAccountAgreement#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

