// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package managementgrouppolicyassignment


type ManagementGroupPolicyAssignmentResourceSelectorsSelectors struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.102.0/docs/resources/management_group_policy_assignment#kind ManagementGroupPolicyAssignment#kind}.
	Kind *string `field:"required" json:"kind" yaml:"kind"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.102.0/docs/resources/management_group_policy_assignment#in ManagementGroupPolicyAssignment#in}.
	In *[]*string `field:"optional" json:"in" yaml:"in"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.102.0/docs/resources/management_group_policy_assignment#not_in ManagementGroupPolicyAssignment#not_in}.
	NotIn *[]*string `field:"optional" json:"notIn" yaml:"notIn"`
}

