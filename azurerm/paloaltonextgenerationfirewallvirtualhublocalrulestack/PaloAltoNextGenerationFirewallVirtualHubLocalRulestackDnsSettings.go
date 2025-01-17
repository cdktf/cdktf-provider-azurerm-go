// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package paloaltonextgenerationfirewallvirtualhublocalrulestack


type PaloAltoNextGenerationFirewallVirtualHubLocalRulestackDnsSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.16.0/docs/resources/palo_alto_next_generation_firewall_virtual_hub_local_rulestack#dns_servers PaloAltoNextGenerationFirewallVirtualHubLocalRulestack#dns_servers}.
	DnsServers *[]*string `field:"optional" json:"dnsServers" yaml:"dnsServers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.16.0/docs/resources/palo_alto_next_generation_firewall_virtual_hub_local_rulestack#use_azure_dns PaloAltoNextGenerationFirewallVirtualHubLocalRulestack#use_azure_dns}.
	UseAzureDns interface{} `field:"optional" json:"useAzureDns" yaml:"useAzureDns"`
}

