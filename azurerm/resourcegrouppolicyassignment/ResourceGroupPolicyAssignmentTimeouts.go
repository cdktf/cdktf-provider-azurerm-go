package resourcegrouppolicyassignment


type ResourceGroupPolicyAssignmentTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/resource_group_policy_assignment#create ResourceGroupPolicyAssignment#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/resource_group_policy_assignment#delete ResourceGroupPolicyAssignment#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/resource_group_policy_assignment#read ResourceGroupPolicyAssignment#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/resource_group_policy_assignment#update ResourceGroupPolicyAssignment#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

