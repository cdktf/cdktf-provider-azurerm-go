package managementgrouppolicyassignment


type ManagementGroupPolicyAssignmentOverrides struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/management_group_policy_assignment#value ManagementGroupPolicyAssignment#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
	// selectors block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/management_group_policy_assignment#selectors ManagementGroupPolicyAssignment#selectors}
	Selectors interface{} `field:"optional" json:"selectors" yaml:"selectors"`
}

