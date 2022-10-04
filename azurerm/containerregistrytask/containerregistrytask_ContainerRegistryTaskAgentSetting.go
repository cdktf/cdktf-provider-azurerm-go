package containerregistrytask


type ContainerRegistryTaskAgentSetting struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_registry_task#cpu ContainerRegistryTask#cpu}.
	Cpu *float64 `field:"required" json:"cpu" yaml:"cpu"`
}

