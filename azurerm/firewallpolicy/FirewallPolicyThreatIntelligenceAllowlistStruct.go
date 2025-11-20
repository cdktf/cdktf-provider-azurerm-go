// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package firewallpolicy


type FirewallPolicyThreatIntelligenceAllowlistStruct struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/firewall_policy#fqdns FirewallPolicy#fqdns}.
	Fqdns *[]*string `field:"optional" json:"fqdns" yaml:"fqdns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/firewall_policy#ip_addresses FirewallPolicy#ip_addresses}.
	IpAddresses *[]*string `field:"optional" json:"ipAddresses" yaml:"ipAddresses"`
}

