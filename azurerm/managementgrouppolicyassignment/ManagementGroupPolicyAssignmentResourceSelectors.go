// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package managementgrouppolicyassignment


type ManagementGroupPolicyAssignmentResourceSelectors struct {
	// selectors block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.83.0/docs/resources/management_group_policy_assignment#selectors ManagementGroupPolicyAssignment#selectors}
	Selectors interface{} `field:"required" json:"selectors" yaml:"selectors"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.83.0/docs/resources/management_group_policy_assignment#name ManagementGroupPolicyAssignment#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

