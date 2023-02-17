package firewallpolicy


type FirewallPolicyThreatIntelligenceAllowlist struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/firewall_policy#fqdns FirewallPolicy#fqdns}.
	Fqdns *[]*string `field:"optional" json:"fqdns" yaml:"fqdns"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/firewall_policy#ip_addresses FirewallPolicy#ip_addresses}.
	IpAddresses *[]*string `field:"optional" json:"ipAddresses" yaml:"ipAddresses"`
}

