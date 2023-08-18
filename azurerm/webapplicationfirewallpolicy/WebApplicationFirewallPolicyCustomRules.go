package webapplicationfirewallpolicy


type WebApplicationFirewallPolicyCustomRules struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/web_application_firewall_policy#action WebApplicationFirewallPolicy#action}.
	Action *string `field:"required" json:"action" yaml:"action"`
	// match_conditions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/web_application_firewall_policy#match_conditions WebApplicationFirewallPolicy#match_conditions}
	MatchConditions interface{} `field:"required" json:"matchConditions" yaml:"matchConditions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/web_application_firewall_policy#priority WebApplicationFirewallPolicy#priority}.
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/web_application_firewall_policy#rule_type WebApplicationFirewallPolicy#rule_type}.
	RuleType *string `field:"required" json:"ruleType" yaml:"ruleType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/web_application_firewall_policy#name WebApplicationFirewallPolicy#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

