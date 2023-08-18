package paloaltonextgenerationfirewallvirtualnetworkpanorama


type PaloAltoNextGenerationFirewallVirtualNetworkPanoramaDestinationNatBackendConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/palo_alto_next_generation_firewall_virtual_network_panorama#port PaloAltoNextGenerationFirewallVirtualNetworkPanorama#port}.
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/palo_alto_next_generation_firewall_virtual_network_panorama#public_ip_address PaloAltoNextGenerationFirewallVirtualNetworkPanorama#public_ip_address}.
	PublicIpAddress *string `field:"required" json:"publicIpAddress" yaml:"publicIpAddress"`
}

