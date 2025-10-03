// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package paloaltonextgenerationfirewallvirtualhubpanorama


type PaloAltoNextGenerationFirewallVirtualHubPanoramaDestinationNat struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.47.0/docs/resources/palo_alto_next_generation_firewall_virtual_hub_panorama#name PaloAltoNextGenerationFirewallVirtualHubPanorama#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.47.0/docs/resources/palo_alto_next_generation_firewall_virtual_hub_panorama#protocol PaloAltoNextGenerationFirewallVirtualHubPanorama#protocol}.
	Protocol *string `field:"required" json:"protocol" yaml:"protocol"`
	// backend_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.47.0/docs/resources/palo_alto_next_generation_firewall_virtual_hub_panorama#backend_config PaloAltoNextGenerationFirewallVirtualHubPanorama#backend_config}
	BackendConfig *PaloAltoNextGenerationFirewallVirtualHubPanoramaDestinationNatBackendConfig `field:"optional" json:"backendConfig" yaml:"backendConfig"`
	// frontend_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.47.0/docs/resources/palo_alto_next_generation_firewall_virtual_hub_panorama#frontend_config PaloAltoNextGenerationFirewallVirtualHubPanorama#frontend_config}
	FrontendConfig *PaloAltoNextGenerationFirewallVirtualHubPanoramaDestinationNatFrontendConfig `field:"optional" json:"frontendConfig" yaml:"frontendConfig"`
}

