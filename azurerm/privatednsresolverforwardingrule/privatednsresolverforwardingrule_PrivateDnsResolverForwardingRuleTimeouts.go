package privatednsresolverforwardingrule


type PrivateDnsResolverForwardingRuleTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/private_dns_resolver_forwarding_rule#create PrivateDnsResolverForwardingRule#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/private_dns_resolver_forwarding_rule#delete PrivateDnsResolverForwardingRule#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/private_dns_resolver_forwarding_rule#read PrivateDnsResolverForwardingRule#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/private_dns_resolver_forwarding_rule#update PrivateDnsResolverForwardingRule#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}
