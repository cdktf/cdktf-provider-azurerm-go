package resourcegrouppolicyremediation


type ResourceGroupPolicyRemediationTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/resource_group_policy_remediation#create ResourceGroupPolicyRemediation#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/resource_group_policy_remediation#delete ResourceGroupPolicyRemediation#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/resource_group_policy_remediation#read ResourceGroupPolicyRemediation#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/resource_group_policy_remediation#update ResourceGroupPolicyRemediation#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}
