package containerregistryscopemap


type ContainerRegistryScopeMapTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_registry_scope_map#create ContainerRegistryScopeMap#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_registry_scope_map#delete ContainerRegistryScopeMap#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_registry_scope_map#read ContainerRegistryScopeMap#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_registry_scope_map#update ContainerRegistryScopeMap#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}
