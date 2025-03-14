// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package paloaltonextgenerationfirewallvirtualnetworkpanorama


type PaloAltoNextGenerationFirewallVirtualNetworkPanoramaDnsSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.23.0/docs/resources/palo_alto_next_generation_firewall_virtual_network_panorama#dns_servers PaloAltoNextGenerationFirewallVirtualNetworkPanorama#dns_servers}.
	DnsServers *[]*string `field:"optional" json:"dnsServers" yaml:"dnsServers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.23.0/docs/resources/palo_alto_next_generation_firewall_virtual_network_panorama#use_azure_dns PaloAltoNextGenerationFirewallVirtualNetworkPanorama#use_azure_dns}.
	UseAzureDns interface{} `field:"optional" json:"useAzureDns" yaml:"useAzureDns"`
}

