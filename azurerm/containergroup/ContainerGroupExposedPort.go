package containergroup


type ContainerGroupExposedPort struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.69.0/docs/resources/container_group#port ContainerGroup#port}.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.69.0/docs/resources/container_group#protocol ContainerGroup#protocol}.
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
}

