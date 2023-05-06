package containerregistry


type ContainerRegistryRetentionPolicy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/container_registry#days ContainerRegistry#days}.
	Days *float64 `field:"optional" json:"days" yaml:"days"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/container_registry#enabled ContainerRegistry#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

