package firewallpolicy


type FirewallPolicyDns struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/firewall_policy#proxy_enabled FirewallPolicy#proxy_enabled}.
	ProxyEnabled interface{} `field:"optional" json:"proxyEnabled" yaml:"proxyEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/firewall_policy#servers FirewallPolicy#servers}.
	Servers *[]*string `field:"optional" json:"servers" yaml:"servers"`
}

