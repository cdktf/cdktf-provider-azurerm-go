package managementgrouppolicyassignment


type ManagementGroupPolicyAssignmentResourceSelectors struct {
	// selectors block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/management_group_policy_assignment#selectors ManagementGroupPolicyAssignment#selectors}
	Selectors interface{} `field:"required" json:"selectors" yaml:"selectors"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/management_group_policy_assignment#name ManagementGroupPolicyAssignment#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

