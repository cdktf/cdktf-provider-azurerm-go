package resourcegrouppolicyexemption


type ResourceGroupPolicyExemptionTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/resource_group_policy_exemption#create ResourceGroupPolicyExemption#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/resource_group_policy_exemption#delete ResourceGroupPolicyExemption#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/resource_group_policy_exemption#read ResourceGroupPolicyExemption#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/resource_group_policy_exemption#update ResourceGroupPolicyExemption#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}
