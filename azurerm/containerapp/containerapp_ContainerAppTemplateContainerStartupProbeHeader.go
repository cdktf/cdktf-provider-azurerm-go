package containerapp


type ContainerAppTemplateContainerStartupProbeHeader struct {
	// The HTTP Header Name.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_app#name ContainerApp#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The HTTP Header value.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_app#value ContainerApp#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

