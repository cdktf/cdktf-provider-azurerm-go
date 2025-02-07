// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resourcepolicyassignment


type ResourcePolicyAssignmentNonComplianceMessage struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/resource_policy_assignment#content ResourcePolicyAssignment#content}.
	Content *string `field:"required" json:"content" yaml:"content"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/resource_policy_assignment#policy_definition_reference_id ResourcePolicyAssignment#policy_definition_reference_id}.
	PolicyDefinitionReferenceId *string `field:"optional" json:"policyDefinitionReferenceId" yaml:"policyDefinitionReferenceId"`
}

