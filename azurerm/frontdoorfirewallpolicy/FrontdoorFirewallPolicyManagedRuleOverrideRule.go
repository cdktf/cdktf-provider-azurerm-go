package frontdoorfirewallpolicy


type FrontdoorFirewallPolicyManagedRuleOverrideRule struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/frontdoor_firewall_policy#action FrontdoorFirewallPolicy#action}.
	Action *string `field:"required" json:"action" yaml:"action"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/frontdoor_firewall_policy#rule_id FrontdoorFirewallPolicy#rule_id}.
	RuleId *string `field:"required" json:"ruleId" yaml:"ruleId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/frontdoor_firewall_policy#enabled FrontdoorFirewallPolicy#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// exclusion block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/frontdoor_firewall_policy#exclusion FrontdoorFirewallPolicy#exclusion}
	Exclusion interface{} `field:"optional" json:"exclusion" yaml:"exclusion"`
}
