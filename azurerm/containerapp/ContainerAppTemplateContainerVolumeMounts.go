package containerapp


type ContainerAppTemplateContainerVolumeMounts struct {
	// The name of the Volume to be mounted in the container.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_app#name ContainerApp#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The path in the container at which to mount this volume.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_app#path ContainerApp#path}
	Path *string `field:"required" json:"path" yaml:"path"`
}
