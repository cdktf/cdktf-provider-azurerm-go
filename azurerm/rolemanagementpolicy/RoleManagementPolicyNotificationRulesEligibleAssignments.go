// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package rolemanagementpolicy


type RoleManagementPolicyNotificationRulesEligibleAssignments struct {
	// admin_notifications block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.42.0/docs/resources/role_management_policy#admin_notifications RoleManagementPolicy#admin_notifications}
	AdminNotifications *RoleManagementPolicyNotificationRulesEligibleAssignmentsAdminNotifications `field:"optional" json:"adminNotifications" yaml:"adminNotifications"`
	// approver_notifications block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.42.0/docs/resources/role_management_policy#approver_notifications RoleManagementPolicy#approver_notifications}
	ApproverNotifications *RoleManagementPolicyNotificationRulesEligibleAssignmentsApproverNotifications `field:"optional" json:"approverNotifications" yaml:"approverNotifications"`
	// assignee_notifications block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.42.0/docs/resources/role_management_policy#assignee_notifications RoleManagementPolicy#assignee_notifications}
	AssigneeNotifications *RoleManagementPolicyNotificationRulesEligibleAssignmentsAssigneeNotifications `field:"optional" json:"assigneeNotifications" yaml:"assigneeNotifications"`
}

