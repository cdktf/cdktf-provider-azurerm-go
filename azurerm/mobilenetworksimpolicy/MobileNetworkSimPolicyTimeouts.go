package mobilenetworksimpolicy


type MobileNetworkSimPolicyTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.66.0/docs/resources/mobile_network_sim_policy#create MobileNetworkSimPolicy#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.66.0/docs/resources/mobile_network_sim_policy#delete MobileNetworkSimPolicy#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.66.0/docs/resources/mobile_network_sim_policy#read MobileNetworkSimPolicy#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.66.0/docs/resources/mobile_network_sim_policy#update MobileNetworkSimPolicy#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

