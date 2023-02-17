package mobilenetworkservice


type MobileNetworkServiceServiceQosPolicyMaximumBitRate struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mobile_network_service#downlink MobileNetworkService#downlink}.
	Downlink *string `field:"required" json:"downlink" yaml:"downlink"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mobile_network_service#uplink MobileNetworkService#uplink}.
	Uplink *string `field:"required" json:"uplink" yaml:"uplink"`
}

