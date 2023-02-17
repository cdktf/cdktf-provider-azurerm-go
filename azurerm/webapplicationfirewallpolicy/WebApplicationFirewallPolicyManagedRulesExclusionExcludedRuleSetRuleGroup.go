package webapplicationfirewallpolicy


type WebApplicationFirewallPolicyManagedRulesExclusionExcludedRuleSetRuleGroup struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/web_application_firewall_policy#rule_group_name WebApplicationFirewallPolicy#rule_group_name}.
	RuleGroupName *string `field:"required" json:"ruleGroupName" yaml:"ruleGroupName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/web_application_firewall_policy#excluded_rules WebApplicationFirewallPolicy#excluded_rules}.
	ExcludedRules *[]*string `field:"optional" json:"excludedRules" yaml:"excludedRules"`
}

