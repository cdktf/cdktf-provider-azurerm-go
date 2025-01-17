// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package rolemanagementpolicy


type RoleManagementPolicyNotificationRules struct {
	// active_assignments block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.16.0/docs/resources/role_management_policy#active_assignments RoleManagementPolicy#active_assignments}
	ActiveAssignments *RoleManagementPolicyNotificationRulesActiveAssignments `field:"optional" json:"activeAssignments" yaml:"activeAssignments"`
	// eligible_activations block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.16.0/docs/resources/role_management_policy#eligible_activations RoleManagementPolicy#eligible_activations}
	EligibleActivations *RoleManagementPolicyNotificationRulesEligibleActivations `field:"optional" json:"eligibleActivations" yaml:"eligibleActivations"`
	// eligible_assignments block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.16.0/docs/resources/role_management_policy#eligible_assignments RoleManagementPolicy#eligible_assignments}
	EligibleAssignments *RoleManagementPolicyNotificationRulesEligibleAssignments `field:"optional" json:"eligibleAssignments" yaml:"eligibleAssignments"`
}

