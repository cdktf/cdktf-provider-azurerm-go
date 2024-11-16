// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package paloaltonextgenerationfirewallvirtualnetworkpanorama


type PaloAltoNextGenerationFirewallVirtualNetworkPanoramaDestinationNat struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/palo_alto_next_generation_firewall_virtual_network_panorama#name PaloAltoNextGenerationFirewallVirtualNetworkPanorama#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/palo_alto_next_generation_firewall_virtual_network_panorama#protocol PaloAltoNextGenerationFirewallVirtualNetworkPanorama#protocol}.
	Protocol *string `field:"required" json:"protocol" yaml:"protocol"`
	// backend_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/palo_alto_next_generation_firewall_virtual_network_panorama#backend_config PaloAltoNextGenerationFirewallVirtualNetworkPanorama#backend_config}
	BackendConfig *PaloAltoNextGenerationFirewallVirtualNetworkPanoramaDestinationNatBackendConfig `field:"optional" json:"backendConfig" yaml:"backendConfig"`
	// frontend_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/palo_alto_next_generation_firewall_virtual_network_panorama#frontend_config PaloAltoNextGenerationFirewallVirtualNetworkPanorama#frontend_config}
	FrontendConfig *PaloAltoNextGenerationFirewallVirtualNetworkPanoramaDestinationNatFrontendConfig `field:"optional" json:"frontendConfig" yaml:"frontendConfig"`
}

