package cdnfrontdoorfirewallpolicy


type CdnFrontdoorFirewallPolicyTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_frontdoor_firewall_policy#create CdnFrontdoorFirewallPolicy#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_frontdoor_firewall_policy#delete CdnFrontdoorFirewallPolicy#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_frontdoor_firewall_policy#read CdnFrontdoorFirewallPolicy#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_frontdoor_firewall_policy#update CdnFrontdoorFirewallPolicy#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}
