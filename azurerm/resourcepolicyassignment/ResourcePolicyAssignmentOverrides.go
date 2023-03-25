package resourcepolicyassignment


type ResourcePolicyAssignmentOverrides struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/resource_policy_assignment#value ResourcePolicyAssignment#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
	// selectors block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/resource_policy_assignment#selectors ResourcePolicyAssignment#selectors}
	Selectors interface{} `field:"optional" json:"selectors" yaml:"selectors"`
}

