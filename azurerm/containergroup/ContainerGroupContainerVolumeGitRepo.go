package containergroup


type ContainerGroupContainerVolumeGitRepo struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_group#url ContainerGroup#url}.
	Url *string `field:"required" json:"url" yaml:"url"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_group#directory ContainerGroup#directory}.
	Directory *string `field:"optional" json:"directory" yaml:"directory"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_group#revision ContainerGroup#revision}.
	Revision *string `field:"optional" json:"revision" yaml:"revision"`
}

