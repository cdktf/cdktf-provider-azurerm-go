package mobilenetworksimpolicy


type MobileNetworkSimPolicySliceDataNetworkSessionAggregateMaximumBitRate struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mobile_network_sim_policy#downlink MobileNetworkSimPolicy#downlink}.
	Downlink *string `field:"required" json:"downlink" yaml:"downlink"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mobile_network_sim_policy#uplink MobileNetworkSimPolicy#uplink}.
	Uplink *string `field:"required" json:"uplink" yaml:"uplink"`
}

