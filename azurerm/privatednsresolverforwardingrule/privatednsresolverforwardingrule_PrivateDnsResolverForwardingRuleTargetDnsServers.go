package privatednsresolverforwardingrule


type PrivateDnsResolverForwardingRuleTargetDnsServers struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/private_dns_resolver_forwarding_rule#ip_address PrivateDnsResolverForwardingRule#ip_address}.
	IpAddress *string `field:"required" json:"ipAddress" yaml:"ipAddress"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/private_dns_resolver_forwarding_rule#port PrivateDnsResolverForwardingRule#port}.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
}
