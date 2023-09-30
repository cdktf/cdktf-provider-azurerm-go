// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package paloaltonextgenerationfirewallvirtualnetworklocalrulestack


type PaloAltoNextGenerationFirewallVirtualNetworkLocalRulestackNetworkProfileVnetConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.75.0/docs/resources/palo_alto_next_generation_firewall_virtual_network_local_rulestack#virtual_network_id PaloAltoNextGenerationFirewallVirtualNetworkLocalRulestack#virtual_network_id}.
	VirtualNetworkId *string `field:"required" json:"virtualNetworkId" yaml:"virtualNetworkId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.75.0/docs/resources/palo_alto_next_generation_firewall_virtual_network_local_rulestack#trusted_subnet_id PaloAltoNextGenerationFirewallVirtualNetworkLocalRulestack#trusted_subnet_id}.
	TrustedSubnetId *string `field:"optional" json:"trustedSubnetId" yaml:"trustedSubnetId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.75.0/docs/resources/palo_alto_next_generation_firewall_virtual_network_local_rulestack#untrusted_subnet_id PaloAltoNextGenerationFirewallVirtualNetworkLocalRulestack#untrusted_subnet_id}.
	UntrustedSubnetId *string `field:"optional" json:"untrustedSubnetId" yaml:"untrustedSubnetId"`
}

