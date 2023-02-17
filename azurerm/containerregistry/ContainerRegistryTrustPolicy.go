package containerregistry


type ContainerRegistryTrustPolicy struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_registry#enabled ContainerRegistry#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

