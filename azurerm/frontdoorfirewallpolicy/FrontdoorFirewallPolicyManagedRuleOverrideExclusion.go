package frontdoorfirewallpolicy


type FrontdoorFirewallPolicyManagedRuleOverrideExclusion struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/frontdoor_firewall_policy#match_variable FrontdoorFirewallPolicy#match_variable}.
	MatchVariable *string `field:"required" json:"matchVariable" yaml:"matchVariable"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/frontdoor_firewall_policy#operator FrontdoorFirewallPolicy#operator}.
	Operator *string `field:"required" json:"operator" yaml:"operator"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/frontdoor_firewall_policy#selector FrontdoorFirewallPolicy#selector}.
	Selector *string `field:"required" json:"selector" yaml:"selector"`
}

