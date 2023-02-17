package webapplicationfirewallpolicy


type WebApplicationFirewallPolicyTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/web_application_firewall_policy#create WebApplicationFirewallPolicy#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/web_application_firewall_policy#delete WebApplicationFirewallPolicy#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/web_application_firewall_policy#read WebApplicationFirewallPolicy#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/web_application_firewall_policy#update WebApplicationFirewallPolicy#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

