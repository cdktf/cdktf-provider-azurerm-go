package mobilenetworkslice


type MobileNetworkSliceTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mobile_network_slice#create MobileNetworkSlice#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mobile_network_slice#delete MobileNetworkSlice#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mobile_network_slice#read MobileNetworkSlice#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mobile_network_slice#update MobileNetworkSlice#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}
