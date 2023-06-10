package containerconnectedregistry


type ContainerConnectedRegistryNotification struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.60.0/docs/resources/container_connected_registry#action ContainerConnectedRegistry#action}.
	Action *string `field:"required" json:"action" yaml:"action"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.60.0/docs/resources/container_connected_registry#name ContainerConnectedRegistry#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.60.0/docs/resources/container_connected_registry#digest ContainerConnectedRegistry#digest}.
	Digest *string `field:"optional" json:"digest" yaml:"digest"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.60.0/docs/resources/container_connected_registry#tag ContainerConnectedRegistry#tag}.
	Tag *string `field:"optional" json:"tag" yaml:"tag"`
}

