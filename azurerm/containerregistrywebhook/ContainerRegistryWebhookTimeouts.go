package containerregistrywebhook


type ContainerRegistryWebhookTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_registry_webhook#create ContainerRegistryWebhook#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_registry_webhook#delete ContainerRegistryWebhook#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_registry_webhook#read ContainerRegistryWebhook#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_registry_webhook#update ContainerRegistryWebhook#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

