package containerregistrytokenpassword


type ContainerRegistryTokenPasswordTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_registry_token_password#create ContainerRegistryTokenPassword#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_registry_token_password#delete ContainerRegistryTokenPassword#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_registry_token_password#read ContainerRegistryTokenPassword#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_registry_token_password#update ContainerRegistryTokenPassword#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

