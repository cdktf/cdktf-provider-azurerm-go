// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package paloaltonextgenerationfirewallvirtualnetworkpanorama


type PaloAltoNextGenerationFirewallVirtualNetworkPanoramaNetworkProfile struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.75.0/docs/resources/palo_alto_next_generation_firewall_virtual_network_panorama#public_ip_address_ids PaloAltoNextGenerationFirewallVirtualNetworkPanorama#public_ip_address_ids}.
	PublicIpAddressIds *[]*string `field:"required" json:"publicIpAddressIds" yaml:"publicIpAddressIds"`
	// vnet_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.75.0/docs/resources/palo_alto_next_generation_firewall_virtual_network_panorama#vnet_configuration PaloAltoNextGenerationFirewallVirtualNetworkPanorama#vnet_configuration}
	VnetConfiguration *PaloAltoNextGenerationFirewallVirtualNetworkPanoramaNetworkProfileVnetConfiguration `field:"required" json:"vnetConfiguration" yaml:"vnetConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.75.0/docs/resources/palo_alto_next_generation_firewall_virtual_network_panorama#egress_nat_ip_address_ids PaloAltoNextGenerationFirewallVirtualNetworkPanorama#egress_nat_ip_address_ids}.
	EgressNatIpAddressIds *[]*string `field:"optional" json:"egressNatIpAddressIds" yaml:"egressNatIpAddressIds"`
}

