// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package paloaltonextgenerationfirewallvirtualnetworklocalrulestack


type PaloAltoNextGenerationFirewallVirtualNetworkLocalRulestackNetworkProfile struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.23.0/docs/resources/palo_alto_next_generation_firewall_virtual_network_local_rulestack#public_ip_address_ids PaloAltoNextGenerationFirewallVirtualNetworkLocalRulestack#public_ip_address_ids}.
	PublicIpAddressIds *[]*string `field:"required" json:"publicIpAddressIds" yaml:"publicIpAddressIds"`
	// vnet_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.23.0/docs/resources/palo_alto_next_generation_firewall_virtual_network_local_rulestack#vnet_configuration PaloAltoNextGenerationFirewallVirtualNetworkLocalRulestack#vnet_configuration}
	VnetConfiguration *PaloAltoNextGenerationFirewallVirtualNetworkLocalRulestackNetworkProfileVnetConfiguration `field:"required" json:"vnetConfiguration" yaml:"vnetConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.23.0/docs/resources/palo_alto_next_generation_firewall_virtual_network_local_rulestack#egress_nat_ip_address_ids PaloAltoNextGenerationFirewallVirtualNetworkLocalRulestack#egress_nat_ip_address_ids}.
	EgressNatIpAddressIds *[]*string `field:"optional" json:"egressNatIpAddressIds" yaml:"egressNatIpAddressIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.23.0/docs/resources/palo_alto_next_generation_firewall_virtual_network_local_rulestack#trusted_address_ranges PaloAltoNextGenerationFirewallVirtualNetworkLocalRulestack#trusted_address_ranges}.
	TrustedAddressRanges *[]*string `field:"optional" json:"trustedAddressRanges" yaml:"trustedAddressRanges"`
}

