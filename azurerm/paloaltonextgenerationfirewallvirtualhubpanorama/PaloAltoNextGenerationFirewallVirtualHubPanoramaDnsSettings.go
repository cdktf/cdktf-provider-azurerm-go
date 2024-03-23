// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package paloaltonextgenerationfirewallvirtualhubpanorama


type PaloAltoNextGenerationFirewallVirtualHubPanoramaDnsSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.97.1/docs/resources/palo_alto_next_generation_firewall_virtual_hub_panorama#dns_servers PaloAltoNextGenerationFirewallVirtualHubPanorama#dns_servers}.
	DnsServers *[]*string `field:"optional" json:"dnsServers" yaml:"dnsServers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.97.1/docs/resources/palo_alto_next_generation_firewall_virtual_hub_panorama#use_azure_dns PaloAltoNextGenerationFirewallVirtualHubPanorama#use_azure_dns}.
	UseAzureDns interface{} `field:"optional" json:"useAzureDns" yaml:"useAzureDns"`
}

