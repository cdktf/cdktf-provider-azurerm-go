package webapplicationfirewallpolicy


type WebApplicationFirewallPolicyManagedRulesExclusionExcludedRuleSet struct {
	// rule_group block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/web_application_firewall_policy#rule_group WebApplicationFirewallPolicy#rule_group}
	RuleGroup interface{} `field:"optional" json:"ruleGroup" yaml:"ruleGroup"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/web_application_firewall_policy#type WebApplicationFirewallPolicy#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/web_application_firewall_policy#version WebApplicationFirewallPolicy#version}.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

