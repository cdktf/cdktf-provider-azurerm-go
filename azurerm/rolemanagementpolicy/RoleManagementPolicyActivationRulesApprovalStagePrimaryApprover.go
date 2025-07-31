// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package rolemanagementpolicy


type RoleManagementPolicyActivationRulesApprovalStagePrimaryApprover struct {
	// The ID of the object to act as an approver.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.0/docs/resources/role_management_policy#object_id RoleManagementPolicy#object_id}
	ObjectId *string `field:"required" json:"objectId" yaml:"objectId"`
	// The type of object acting as an approver.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.0/docs/resources/role_management_policy#type RoleManagementPolicy#type}
	Type *string `field:"required" json:"type" yaml:"type"`
}

