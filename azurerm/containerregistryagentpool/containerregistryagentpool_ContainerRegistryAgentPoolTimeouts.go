package containerregistryagentpool


type ContainerRegistryAgentPoolTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_registry_agent_pool#create ContainerRegistryAgentPool#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_registry_agent_pool#delete ContainerRegistryAgentPool#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_registry_agent_pool#read ContainerRegistryAgentPool#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_registry_agent_pool#update ContainerRegistryAgentPool#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

