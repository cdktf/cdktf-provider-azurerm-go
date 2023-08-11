package mobilenetworksim


type MobileNetworkSimStaticIpConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.69.0/docs/resources/mobile_network_sim#attached_data_network_id MobileNetworkSim#attached_data_network_id}.
	AttachedDataNetworkId *string `field:"required" json:"attachedDataNetworkId" yaml:"attachedDataNetworkId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.69.0/docs/resources/mobile_network_sim#slice_id MobileNetworkSim#slice_id}.
	SliceId *string `field:"required" json:"sliceId" yaml:"sliceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.69.0/docs/resources/mobile_network_sim#static_ipv4_address MobileNetworkSim#static_ipv4_address}.
	StaticIpv4Address *string `field:"optional" json:"staticIpv4Address" yaml:"staticIpv4Address"`
}

