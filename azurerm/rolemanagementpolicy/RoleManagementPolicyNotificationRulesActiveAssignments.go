// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package rolemanagementpolicy


type RoleManagementPolicyNotificationRulesActiveAssignments struct {
	// admin_notifications block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.35.0/docs/resources/role_management_policy#admin_notifications RoleManagementPolicy#admin_notifications}
	AdminNotifications *RoleManagementPolicyNotificationRulesActiveAssignmentsAdminNotifications `field:"optional" json:"adminNotifications" yaml:"adminNotifications"`
	// approver_notifications block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.35.0/docs/resources/role_management_policy#approver_notifications RoleManagementPolicy#approver_notifications}
	ApproverNotifications *RoleManagementPolicyNotificationRulesActiveAssignmentsApproverNotifications `field:"optional" json:"approverNotifications" yaml:"approverNotifications"`
	// assignee_notifications block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.35.0/docs/resources/role_management_policy#assignee_notifications RoleManagementPolicy#assignee_notifications}
	AssigneeNotifications *RoleManagementPolicyNotificationRulesActiveAssignmentsAssigneeNotifications `field:"optional" json:"assigneeNotifications" yaml:"assigneeNotifications"`
}

