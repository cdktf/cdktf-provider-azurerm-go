package containerregistrytask


type ContainerRegistryTaskRegistryCredentialSource struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_registry_task#login_mode ContainerRegistryTask#login_mode}.
	LoginMode *string `field:"required" json:"loginMode" yaml:"loginMode"`
}

