package containerregistry


type ContainerRegistryNetworkRuleSet struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_registry#default_action ContainerRegistry#default_action}.
	DefaultAction *string `field:"optional" json:"defaultAction" yaml:"defaultAction"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_registry#ip_rule ContainerRegistry#ip_rule}.
	IpRule interface{} `field:"optional" json:"ipRule" yaml:"ipRule"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_registry#virtual_network ContainerRegistry#virtual_network}.
	VirtualNetwork interface{} `field:"optional" json:"virtualNetwork" yaml:"virtualNetwork"`
}
