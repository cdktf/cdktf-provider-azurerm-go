package managementgrouppolicyassignment


type ManagementGroupPolicyAssignmentNonComplianceMessage struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/management_group_policy_assignment#content ManagementGroupPolicyAssignment#content}.
	Content *string `field:"required" json:"content" yaml:"content"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/management_group_policy_assignment#policy_definition_reference_id ManagementGroupPolicyAssignment#policy_definition_reference_id}.
	PolicyDefinitionReferenceId *string `field:"optional" json:"policyDefinitionReferenceId" yaml:"policyDefinitionReferenceId"`
}

