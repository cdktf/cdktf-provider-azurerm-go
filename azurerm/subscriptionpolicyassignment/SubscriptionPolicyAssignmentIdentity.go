package subscriptionpolicyassignment


type SubscriptionPolicyAssignmentIdentity struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/subscription_policy_assignment#type SubscriptionPolicyAssignment#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/subscription_policy_assignment#identity_ids SubscriptionPolicyAssignment#identity_ids}.
	IdentityIds *[]*string `field:"optional" json:"identityIds" yaml:"identityIds"`
}

