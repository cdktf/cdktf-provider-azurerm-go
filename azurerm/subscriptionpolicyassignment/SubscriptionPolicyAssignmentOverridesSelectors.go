package subscriptionpolicyassignment


type SubscriptionPolicyAssignmentOverridesSelectors struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/subscription_policy_assignment#in SubscriptionPolicyAssignment#in}.
	In *[]*string `field:"optional" json:"in" yaml:"in"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/subscription_policy_assignment#not_in SubscriptionPolicyAssignment#not_in}.
	NotIn *[]*string `field:"optional" json:"notIn" yaml:"notIn"`
}

