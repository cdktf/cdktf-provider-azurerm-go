package containerregistrytask


type ContainerRegistryTaskTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_registry_task#create ContainerRegistryTask#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_registry_task#delete ContainerRegistryTask#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_registry_task#read ContainerRegistryTask#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_registry_task#update ContainerRegistryTask#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}
