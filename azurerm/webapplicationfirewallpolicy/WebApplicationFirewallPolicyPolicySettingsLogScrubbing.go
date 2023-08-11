package webapplicationfirewallpolicy


type WebApplicationFirewallPolicyPolicySettingsLogScrubbing struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.69.0/docs/resources/web_application_firewall_policy#enabled WebApplicationFirewallPolicy#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.69.0/docs/resources/web_application_firewall_policy#rule WebApplicationFirewallPolicy#rule}
	Rule interface{} `field:"optional" json:"rule" yaml:"rule"`
}

