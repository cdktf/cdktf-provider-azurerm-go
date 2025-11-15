// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package rolemanagementpolicy


type RoleManagementPolicyActivationRulesApprovalStage struct {
	// primary_approver block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.53.0/docs/resources/role_management_policy#primary_approver RoleManagementPolicy#primary_approver}
	PrimaryApprover interface{} `field:"required" json:"primaryApprover" yaml:"primaryApprover"`
}

