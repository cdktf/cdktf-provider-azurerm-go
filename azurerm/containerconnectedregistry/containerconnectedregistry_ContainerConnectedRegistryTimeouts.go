package containerconnectedregistry


type ContainerConnectedRegistryTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_connected_registry#create ContainerConnectedRegistry#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_connected_registry#delete ContainerConnectedRegistry#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_connected_registry#read ContainerConnectedRegistry#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_connected_registry#update ContainerConnectedRegistry#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

