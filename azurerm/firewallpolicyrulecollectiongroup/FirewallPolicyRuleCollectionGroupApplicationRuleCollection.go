package firewallpolicyrulecollectiongroup


type FirewallPolicyRuleCollectionGroupApplicationRuleCollection struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/firewall_policy_rule_collection_group#action FirewallPolicyRuleCollectionGroup#action}.
	Action *string `field:"required" json:"action" yaml:"action"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/firewall_policy_rule_collection_group#name FirewallPolicyRuleCollectionGroup#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/firewall_policy_rule_collection_group#priority FirewallPolicyRuleCollectionGroup#priority}.
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
	// rule block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/firewall_policy_rule_collection_group#rule FirewallPolicyRuleCollectionGroup#rule}
	Rule interface{} `field:"required" json:"rule" yaml:"rule"`
}

