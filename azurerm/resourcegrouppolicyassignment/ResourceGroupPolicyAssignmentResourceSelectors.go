// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resourcegrouppolicyassignment


type ResourceGroupPolicyAssignmentResourceSelectors struct {
	// selectors block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/resources/resource_group_policy_assignment#selectors ResourceGroupPolicyAssignment#selectors}
	Selectors interface{} `field:"required" json:"selectors" yaml:"selectors"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/resources/resource_group_policy_assignment#name ResourceGroupPolicyAssignment#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

