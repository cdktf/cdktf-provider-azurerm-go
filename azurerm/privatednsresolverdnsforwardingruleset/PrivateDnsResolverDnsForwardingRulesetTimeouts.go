package privatednsresolverdnsforwardingruleset


type PrivateDnsResolverDnsForwardingRulesetTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/private_dns_resolver_dns_forwarding_ruleset#create PrivateDnsResolverDnsForwardingRuleset#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/private_dns_resolver_dns_forwarding_ruleset#delete PrivateDnsResolverDnsForwardingRuleset#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/private_dns_resolver_dns_forwarding_ruleset#read PrivateDnsResolverDnsForwardingRuleset#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/private_dns_resolver_dns_forwarding_ruleset#update PrivateDnsResolverDnsForwardingRuleset#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}
