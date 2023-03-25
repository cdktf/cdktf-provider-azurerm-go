package resourcepolicyassignment


type ResourcePolicyAssignmentResourceSelectors struct {
	// selectors block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/resource_policy_assignment#selectors ResourcePolicyAssignment#selectors}
	Selectors interface{} `field:"required" json:"selectors" yaml:"selectors"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/resource_policy_assignment#name ResourcePolicyAssignment#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

