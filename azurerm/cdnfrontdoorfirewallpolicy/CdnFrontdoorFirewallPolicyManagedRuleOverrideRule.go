package cdnfrontdoorfirewallpolicy


type CdnFrontdoorFirewallPolicyManagedRuleOverrideRule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.69.0/docs/resources/cdn_frontdoor_firewall_policy#action CdnFrontdoorFirewallPolicy#action}.
	Action *string `field:"required" json:"action" yaml:"action"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.69.0/docs/resources/cdn_frontdoor_firewall_policy#rule_id CdnFrontdoorFirewallPolicy#rule_id}.
	RuleId *string `field:"required" json:"ruleId" yaml:"ruleId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.69.0/docs/resources/cdn_frontdoor_firewall_policy#enabled CdnFrontdoorFirewallPolicy#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// exclusion block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.69.0/docs/resources/cdn_frontdoor_firewall_policy#exclusion CdnFrontdoorFirewallPolicy#exclusion}
	Exclusion interface{} `field:"optional" json:"exclusion" yaml:"exclusion"`
}

