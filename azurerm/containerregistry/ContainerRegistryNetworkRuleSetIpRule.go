package containerregistry


type ContainerRegistryNetworkRuleSetIpRule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.59.0/docs/resources/container_registry#action ContainerRegistry#action}.
	Action *string `field:"optional" json:"action" yaml:"action"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.59.0/docs/resources/container_registry#ip_range ContainerRegistry#ip_range}.
	IpRange *string `field:"optional" json:"ipRange" yaml:"ipRange"`
}

