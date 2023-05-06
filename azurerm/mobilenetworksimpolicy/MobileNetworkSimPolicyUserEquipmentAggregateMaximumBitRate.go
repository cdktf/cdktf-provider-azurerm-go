package mobilenetworksimpolicy


type MobileNetworkSimPolicyUserEquipmentAggregateMaximumBitRate struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/mobile_network_sim_policy#downlink MobileNetworkSimPolicy#downlink}.
	Downlink *string `field:"required" json:"downlink" yaml:"downlink"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/mobile_network_sim_policy#uplink MobileNetworkSimPolicy#uplink}.
	Uplink *string `field:"required" json:"uplink" yaml:"uplink"`
}

