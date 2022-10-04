package virtualhub


type VirtualHubRoute struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_hub#address_prefixes VirtualHub#address_prefixes}.
	AddressPrefixes *[]*string `field:"required" json:"addressPrefixes" yaml:"addressPrefixes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_hub#next_hop_ip_address VirtualHub#next_hop_ip_address}.
	NextHopIpAddress *string `field:"required" json:"nextHopIpAddress" yaml:"nextHopIpAddress"`
}

