package containerregistrytask


type ContainerRegistryTaskRegistryCredential struct {
	// custom block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_registry_task#custom ContainerRegistryTask#custom}
	Custom interface{} `field:"optional" json:"custom" yaml:"custom"`
	// source block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_registry_task#source ContainerRegistryTask#source}
	Source *ContainerRegistryTaskRegistryCredentialSource `field:"optional" json:"source" yaml:"source"`
}

