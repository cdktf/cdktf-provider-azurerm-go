// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package paloaltonextgenerationfirewallvirtualnetworklocalrulestack


type PaloAltoNextGenerationFirewallVirtualNetworkLocalRulestackDestinationNatFrontendConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/palo_alto_next_generation_firewall_virtual_network_local_rulestack#port PaloAltoNextGenerationFirewallVirtualNetworkLocalRulestack#port}.
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/palo_alto_next_generation_firewall_virtual_network_local_rulestack#public_ip_address_id PaloAltoNextGenerationFirewallVirtualNetworkLocalRulestack#public_ip_address_id}.
	PublicIpAddressId *string `field:"required" json:"publicIpAddressId" yaml:"publicIpAddressId"`
}

