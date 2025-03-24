// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package paloaltonextgenerationfirewallvirtualnetworkpanorama


type PaloAltoNextGenerationFirewallVirtualNetworkPanoramaNetworkProfileVnetConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.24.0/docs/resources/palo_alto_next_generation_firewall_virtual_network_panorama#virtual_network_id PaloAltoNextGenerationFirewallVirtualNetworkPanorama#virtual_network_id}.
	VirtualNetworkId *string `field:"required" json:"virtualNetworkId" yaml:"virtualNetworkId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.24.0/docs/resources/palo_alto_next_generation_firewall_virtual_network_panorama#trusted_subnet_id PaloAltoNextGenerationFirewallVirtualNetworkPanorama#trusted_subnet_id}.
	TrustedSubnetId *string `field:"optional" json:"trustedSubnetId" yaml:"trustedSubnetId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.24.0/docs/resources/palo_alto_next_generation_firewall_virtual_network_panorama#untrusted_subnet_id PaloAltoNextGenerationFirewallVirtualNetworkPanorama#untrusted_subnet_id}.
	UntrustedSubnetId *string `field:"optional" json:"untrustedSubnetId" yaml:"untrustedSubnetId"`
}

