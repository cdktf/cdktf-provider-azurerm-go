package resourcegrouppolicyassignment


type ResourceGroupPolicyAssignmentOverrides struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.68.0/docs/resources/resource_group_policy_assignment#value ResourceGroupPolicyAssignment#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
	// selectors block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.68.0/docs/resources/resource_group_policy_assignment#selectors ResourceGroupPolicyAssignment#selectors}
	Selectors interface{} `field:"optional" json:"selectors" yaml:"selectors"`
}

