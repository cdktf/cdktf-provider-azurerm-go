package containerapp


type ContainerAppTemplateContainerReadinessProbeHeader struct {
	// The HTTP Header Name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.57.0/docs/resources/container_app#name ContainerApp#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The HTTP Header value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.57.0/docs/resources/container_app#value ContainerApp#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

