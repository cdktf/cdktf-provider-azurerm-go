package containergroup


type ContainerGroupContainerGpu struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.0/docs/resources/container_group#count ContainerGroup#count}.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.0/docs/resources/container_group#sku ContainerGroup#sku}.
	Sku *string `field:"optional" json:"sku" yaml:"sku"`
}

