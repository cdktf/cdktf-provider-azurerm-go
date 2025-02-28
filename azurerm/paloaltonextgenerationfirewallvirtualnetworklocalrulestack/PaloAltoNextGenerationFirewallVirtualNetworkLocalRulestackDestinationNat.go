// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package paloaltonextgenerationfirewallvirtualnetworklocalrulestack


type PaloAltoNextGenerationFirewallVirtualNetworkLocalRulestackDestinationNat struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.0/docs/resources/palo_alto_next_generation_firewall_virtual_network_local_rulestack#name PaloAltoNextGenerationFirewallVirtualNetworkLocalRulestack#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.0/docs/resources/palo_alto_next_generation_firewall_virtual_network_local_rulestack#protocol PaloAltoNextGenerationFirewallVirtualNetworkLocalRulestack#protocol}.
	Protocol *string `field:"required" json:"protocol" yaml:"protocol"`
	// backend_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.0/docs/resources/palo_alto_next_generation_firewall_virtual_network_local_rulestack#backend_config PaloAltoNextGenerationFirewallVirtualNetworkLocalRulestack#backend_config}
	BackendConfig *PaloAltoNextGenerationFirewallVirtualNetworkLocalRulestackDestinationNatBackendConfig `field:"optional" json:"backendConfig" yaml:"backendConfig"`
	// frontend_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.0/docs/resources/palo_alto_next_generation_firewall_virtual_network_local_rulestack#frontend_config PaloAltoNextGenerationFirewallVirtualNetworkLocalRulestack#frontend_config}
	FrontendConfig *PaloAltoNextGenerationFirewallVirtualNetworkLocalRulestackDestinationNatFrontendConfig `field:"optional" json:"frontendConfig" yaml:"frontendConfig"`
}

