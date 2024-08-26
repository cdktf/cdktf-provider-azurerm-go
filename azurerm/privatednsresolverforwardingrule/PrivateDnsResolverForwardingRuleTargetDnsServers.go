// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package privatednsresolverforwardingrule


type PrivateDnsResolverForwardingRuleTargetDnsServers struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.0.1/docs/resources/private_dns_resolver_forwarding_rule#ip_address PrivateDnsResolverForwardingRule#ip_address}.
	IpAddress *string `field:"required" json:"ipAddress" yaml:"ipAddress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.0.1/docs/resources/private_dns_resolver_forwarding_rule#port PrivateDnsResolverForwardingRule#port}.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
}

