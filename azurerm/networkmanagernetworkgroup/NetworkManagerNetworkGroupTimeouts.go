package networkmanagernetworkgroup


type NetworkManagerNetworkGroupTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/network_manager_network_group#create NetworkManagerNetworkGroup#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/network_manager_network_group#delete NetworkManagerNetworkGroup#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/network_manager_network_group#read NetworkManagerNetworkGroup#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/network_manager_network_group#update NetworkManagerNetworkGroup#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

