package containerregistry


type ContainerRegistryNetworkRuleSetIpRule struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_registry#action ContainerRegistry#action}.
	Action *string `field:"optional" json:"action" yaml:"action"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_registry#ip_range ContainerRegistry#ip_range}.
	IpRange *string `field:"optional" json:"ipRange" yaml:"ipRange"`
}

