// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package managementgrouppolicyassignment


type ManagementGroupPolicyAssignmentTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/management_group_policy_assignment#create ManagementGroupPolicyAssignment#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/management_group_policy_assignment#delete ManagementGroupPolicyAssignment#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/management_group_policy_assignment#read ManagementGroupPolicyAssignment#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/management_group_policy_assignment#update ManagementGroupPolicyAssignment#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

