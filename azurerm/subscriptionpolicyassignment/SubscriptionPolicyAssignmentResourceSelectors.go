package subscriptionpolicyassignment


type SubscriptionPolicyAssignmentResourceSelectors struct {
	// selectors block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/subscription_policy_assignment#selectors SubscriptionPolicyAssignment#selectors}
	Selectors interface{} `field:"required" json:"selectors" yaml:"selectors"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/subscription_policy_assignment#name SubscriptionPolicyAssignment#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

