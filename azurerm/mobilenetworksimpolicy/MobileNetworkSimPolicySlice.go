package mobilenetworksimpolicy


type MobileNetworkSimPolicySlice struct {
	// data_network block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/mobile_network_sim_policy#data_network MobileNetworkSimPolicy#data_network}
	DataNetwork interface{} `field:"required" json:"dataNetwork" yaml:"dataNetwork"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/mobile_network_sim_policy#default_data_network_id MobileNetworkSimPolicy#default_data_network_id}.
	DefaultDataNetworkId *string `field:"required" json:"defaultDataNetworkId" yaml:"defaultDataNetworkId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/mobile_network_sim_policy#slice_id MobileNetworkSimPolicy#slice_id}.
	SliceId *string `field:"required" json:"sliceId" yaml:"sliceId"`
}

