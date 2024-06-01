// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lighthousedefinition


type LighthouseDefinitionEligibleAuthorizationJustInTimeAccessPolicy struct {
	// approver block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.106.1/docs/resources/lighthouse_definition#approver LighthouseDefinition#approver}
	Approver interface{} `field:"optional" json:"approver" yaml:"approver"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.106.1/docs/resources/lighthouse_definition#maximum_activation_duration LighthouseDefinition#maximum_activation_duration}.
	MaximumActivationDuration *string `field:"optional" json:"maximumActivationDuration" yaml:"maximumActivationDuration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.106.1/docs/resources/lighthouse_definition#multi_factor_auth_provider LighthouseDefinition#multi_factor_auth_provider}.
	MultiFactorAuthProvider *string `field:"optional" json:"multiFactorAuthProvider" yaml:"multiFactorAuthProvider"`
}

