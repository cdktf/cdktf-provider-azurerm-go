package paloaltonextgenerationfirewallvirtualhublocalrulestack


type PaloAltoNextGenerationFirewallVirtualHubLocalRulestackNetworkProfile struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/palo_alto_next_generation_firewall_virtual_hub_local_rulestack#network_virtual_appliance_id PaloAltoNextGenerationFirewallVirtualHubLocalRulestack#network_virtual_appliance_id}.
	NetworkVirtualApplianceId *string `field:"required" json:"networkVirtualApplianceId" yaml:"networkVirtualApplianceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/palo_alto_next_generation_firewall_virtual_hub_local_rulestack#public_ip_address_ids PaloAltoNextGenerationFirewallVirtualHubLocalRulestack#public_ip_address_ids}.
	PublicIpAddressIds *[]*string `field:"required" json:"publicIpAddressIds" yaml:"publicIpAddressIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/palo_alto_next_generation_firewall_virtual_hub_local_rulestack#virtual_hub_id PaloAltoNextGenerationFirewallVirtualHubLocalRulestack#virtual_hub_id}.
	VirtualHubId *string `field:"required" json:"virtualHubId" yaml:"virtualHubId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/palo_alto_next_generation_firewall_virtual_hub_local_rulestack#egress_nat_ip_address_ids PaloAltoNextGenerationFirewallVirtualHubLocalRulestack#egress_nat_ip_address_ids}.
	EgressNatIpAddressIds *[]*string `field:"optional" json:"egressNatIpAddressIds" yaml:"egressNatIpAddressIds"`
}

