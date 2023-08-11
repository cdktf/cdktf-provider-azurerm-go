package paloaltonextgenerationfirewallvirtualhubpanorama


type PaloAltoNextGenerationFirewallVirtualHubPanoramaDestinationNatFrontendConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.69.0/docs/resources/palo_alto_next_generation_firewall_virtual_hub_panorama#port PaloAltoNextGenerationFirewallVirtualHubPanorama#port}.
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.69.0/docs/resources/palo_alto_next_generation_firewall_virtual_hub_panorama#public_ip_address_id PaloAltoNextGenerationFirewallVirtualHubPanorama#public_ip_address_id}.
	PublicIpAddressId *string `field:"required" json:"publicIpAddressId" yaml:"publicIpAddressId"`
}

