// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lighthousedefinition


type LighthouseDefinitionEligibleAuthorizationJustInTimeAccessPolicyApprover struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.101.0/docs/resources/lighthouse_definition#principal_id LighthouseDefinition#principal_id}.
	PrincipalId *string `field:"required" json:"principalId" yaml:"principalId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.101.0/docs/resources/lighthouse_definition#principal_display_name LighthouseDefinition#principal_display_name}.
	PrincipalDisplayName *string `field:"optional" json:"principalDisplayName" yaml:"principalDisplayName"`
}

