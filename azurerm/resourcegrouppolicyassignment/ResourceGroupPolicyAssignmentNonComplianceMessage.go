// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resourcegrouppolicyassignment


type ResourceGroupPolicyAssignmentNonComplianceMessage struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/resource_group_policy_assignment#content ResourceGroupPolicyAssignment#content}.
	Content *string `field:"required" json:"content" yaml:"content"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/resource_group_policy_assignment#policy_definition_reference_id ResourceGroupPolicyAssignment#policy_definition_reference_id}.
	PolicyDefinitionReferenceId *string `field:"optional" json:"policyDefinitionReferenceId" yaml:"policyDefinitionReferenceId"`
}

