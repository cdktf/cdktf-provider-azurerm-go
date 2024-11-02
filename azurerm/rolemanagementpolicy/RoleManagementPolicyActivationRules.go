// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package rolemanagementpolicy


type RoleManagementPolicyActivationRules struct {
	// approval_stage block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/role_management_policy#approval_stage RoleManagementPolicy#approval_stage}
	ApprovalStage *RoleManagementPolicyActivationRulesApprovalStage `field:"optional" json:"approvalStage" yaml:"approvalStage"`
	// The time after which the an activation can be valid for.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/role_management_policy#maximum_duration RoleManagementPolicy#maximum_duration}
	MaximumDuration *string `field:"optional" json:"maximumDuration" yaml:"maximumDuration"`
	// Whether an approval is required for activation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/role_management_policy#require_approval RoleManagementPolicy#require_approval}
	RequireApproval interface{} `field:"optional" json:"requireApproval" yaml:"requireApproval"`
	// Whether a conditional access context is required during activation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/role_management_policy#required_conditional_access_authentication_context RoleManagementPolicy#required_conditional_access_authentication_context}
	RequiredConditionalAccessAuthenticationContext *string `field:"optional" json:"requiredConditionalAccessAuthenticationContext" yaml:"requiredConditionalAccessAuthenticationContext"`
	// Whether a justification is required during activation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/role_management_policy#require_justification RoleManagementPolicy#require_justification}
	RequireJustification interface{} `field:"optional" json:"requireJustification" yaml:"requireJustification"`
	// Whether multi-factor authentication is required during activation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/role_management_policy#require_multifactor_authentication RoleManagementPolicy#require_multifactor_authentication}
	RequireMultifactorAuthentication interface{} `field:"optional" json:"requireMultifactorAuthentication" yaml:"requireMultifactorAuthentication"`
	// Whether ticket information is required during activation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/role_management_policy#require_ticket_info RoleManagementPolicy#require_ticket_info}
	RequireTicketInfo interface{} `field:"optional" json:"requireTicketInfo" yaml:"requireTicketInfo"`
}

