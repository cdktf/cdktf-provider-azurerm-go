package webapplicationfirewallpolicy


type WebApplicationFirewallPolicyManagedRules struct {
	// managed_rule_set block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/web_application_firewall_policy#managed_rule_set WebApplicationFirewallPolicy#managed_rule_set}
	ManagedRuleSet interface{} `field:"required" json:"managedRuleSet" yaml:"managedRuleSet"`
	// exclusion block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/web_application_firewall_policy#exclusion WebApplicationFirewallPolicy#exclusion}
	Exclusion interface{} `field:"optional" json:"exclusion" yaml:"exclusion"`
}
