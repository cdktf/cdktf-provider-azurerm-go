package mobilenetworkservice


type MobileNetworkServicePccRuleQosPolicyGuaranteedBitRate struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.58.0/docs/resources/mobile_network_service#downlink MobileNetworkService#downlink}.
	Downlink *string `field:"required" json:"downlink" yaml:"downlink"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.58.0/docs/resources/mobile_network_service#uplink MobileNetworkService#uplink}.
	Uplink *string `field:"required" json:"uplink" yaml:"uplink"`
}

