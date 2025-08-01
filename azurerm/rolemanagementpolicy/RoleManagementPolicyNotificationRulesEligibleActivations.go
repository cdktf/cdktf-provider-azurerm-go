// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package rolemanagementpolicy


type RoleManagementPolicyNotificationRulesEligibleActivations struct {
	// admin_notifications block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/role_management_policy#admin_notifications RoleManagementPolicy#admin_notifications}
	AdminNotifications *RoleManagementPolicyNotificationRulesEligibleActivationsAdminNotifications `field:"optional" json:"adminNotifications" yaml:"adminNotifications"`
	// approver_notifications block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/role_management_policy#approver_notifications RoleManagementPolicy#approver_notifications}
	ApproverNotifications *RoleManagementPolicyNotificationRulesEligibleActivationsApproverNotifications `field:"optional" json:"approverNotifications" yaml:"approverNotifications"`
	// assignee_notifications block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/role_management_policy#assignee_notifications RoleManagementPolicy#assignee_notifications}
	AssigneeNotifications *RoleManagementPolicyNotificationRulesEligibleActivationsAssigneeNotifications `field:"optional" json:"assigneeNotifications" yaml:"assigneeNotifications"`
}

