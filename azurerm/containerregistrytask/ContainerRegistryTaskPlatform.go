package containerregistrytask


type ContainerRegistryTaskPlatform struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_registry_task#os ContainerRegistryTask#os}.
	Os *string `field:"required" json:"os" yaml:"os"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_registry_task#architecture ContainerRegistryTask#architecture}.
	Architecture *string `field:"optional" json:"architecture" yaml:"architecture"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_registry_task#variant ContainerRegistryTask#variant}.
	Variant *string `field:"optional" json:"variant" yaml:"variant"`
}

