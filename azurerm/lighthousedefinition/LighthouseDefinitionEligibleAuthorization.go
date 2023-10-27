// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lighthousedefinition


type LighthouseDefinitionEligibleAuthorization struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.78.0/docs/resources/lighthouse_definition#principal_id LighthouseDefinition#principal_id}.
	PrincipalId *string `field:"required" json:"principalId" yaml:"principalId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.78.0/docs/resources/lighthouse_definition#role_definition_id LighthouseDefinition#role_definition_id}.
	RoleDefinitionId *string `field:"required" json:"roleDefinitionId" yaml:"roleDefinitionId"`
	// just_in_time_access_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.78.0/docs/resources/lighthouse_definition#just_in_time_access_policy LighthouseDefinition#just_in_time_access_policy}
	JustInTimeAccessPolicy *LighthouseDefinitionEligibleAuthorizationJustInTimeAccessPolicy `field:"optional" json:"justInTimeAccessPolicy" yaml:"justInTimeAccessPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.78.0/docs/resources/lighthouse_definition#principal_display_name LighthouseDefinition#principal_display_name}.
	PrincipalDisplayName *string `field:"optional" json:"principalDisplayName" yaml:"principalDisplayName"`
}

