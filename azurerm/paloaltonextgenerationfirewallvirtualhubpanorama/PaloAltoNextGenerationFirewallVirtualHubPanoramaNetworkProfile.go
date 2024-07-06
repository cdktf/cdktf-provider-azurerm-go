// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package paloaltonextgenerationfirewallvirtualhubpanorama


type PaloAltoNextGenerationFirewallVirtualHubPanoramaNetworkProfile struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.111.0/docs/resources/palo_alto_next_generation_firewall_virtual_hub_panorama#network_virtual_appliance_id PaloAltoNextGenerationFirewallVirtualHubPanorama#network_virtual_appliance_id}.
	NetworkVirtualApplianceId *string `field:"required" json:"networkVirtualApplianceId" yaml:"networkVirtualApplianceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.111.0/docs/resources/palo_alto_next_generation_firewall_virtual_hub_panorama#public_ip_address_ids PaloAltoNextGenerationFirewallVirtualHubPanorama#public_ip_address_ids}.
	PublicIpAddressIds *[]*string `field:"required" json:"publicIpAddressIds" yaml:"publicIpAddressIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.111.0/docs/resources/palo_alto_next_generation_firewall_virtual_hub_panorama#virtual_hub_id PaloAltoNextGenerationFirewallVirtualHubPanorama#virtual_hub_id}.
	VirtualHubId *string `field:"required" json:"virtualHubId" yaml:"virtualHubId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.111.0/docs/resources/palo_alto_next_generation_firewall_virtual_hub_panorama#egress_nat_ip_address_ids PaloAltoNextGenerationFirewallVirtualHubPanorama#egress_nat_ip_address_ids}.
	EgressNatIpAddressIds *[]*string `field:"optional" json:"egressNatIpAddressIds" yaml:"egressNatIpAddressIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.111.0/docs/resources/palo_alto_next_generation_firewall_virtual_hub_panorama#trusted_address_ranges PaloAltoNextGenerationFirewallVirtualHubPanorama#trusted_address_ranges}.
	TrustedAddressRanges *[]*string `field:"optional" json:"trustedAddressRanges" yaml:"trustedAddressRanges"`
}

