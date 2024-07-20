// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package paloaltonextgenerationfirewallvirtualnetworklocalrulestack


type PaloAltoNextGenerationFirewallVirtualNetworkLocalRulestackDnsSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.113.0/docs/resources/palo_alto_next_generation_firewall_virtual_network_local_rulestack#dns_servers PaloAltoNextGenerationFirewallVirtualNetworkLocalRulestack#dns_servers}.
	DnsServers *[]*string `field:"optional" json:"dnsServers" yaml:"dnsServers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.113.0/docs/resources/palo_alto_next_generation_firewall_virtual_network_local_rulestack#use_azure_dns PaloAltoNextGenerationFirewallVirtualNetworkLocalRulestack#use_azure_dns}.
	UseAzureDns interface{} `field:"optional" json:"useAzureDns" yaml:"useAzureDns"`
}

