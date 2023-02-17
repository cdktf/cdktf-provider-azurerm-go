package managementgrouppolicyassignment


type ManagementGroupPolicyAssignmentIdentity struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/management_group_policy_assignment#type ManagementGroupPolicyAssignment#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/management_group_policy_assignment#identity_ids ManagementGroupPolicyAssignment#identity_ids}.
	IdentityIds *[]*string `field:"optional" json:"identityIds" yaml:"identityIds"`
}

