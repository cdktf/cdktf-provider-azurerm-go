// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package managementgrouppolicyassignment


type ManagementGroupPolicyAssignmentOverrides struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.87.0/docs/resources/management_group_policy_assignment#value ManagementGroupPolicyAssignment#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
	// selectors block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.87.0/docs/resources/management_group_policy_assignment#selectors ManagementGroupPolicyAssignment#selectors}
	Selectors interface{} `field:"optional" json:"selectors" yaml:"selectors"`
}

