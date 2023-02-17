package containerregistry


type ContainerRegistryRetentionPolicy struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_registry#days ContainerRegistry#days}.
	Days *float64 `field:"optional" json:"days" yaml:"days"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_registry#enabled ContainerRegistry#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

