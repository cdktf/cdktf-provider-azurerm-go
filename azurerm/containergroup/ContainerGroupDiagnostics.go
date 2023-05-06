package containergroup


type ContainerGroupDiagnostics struct {
	// log_analytics block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/container_group#log_analytics ContainerGroup#log_analytics}
	LogAnalytics *ContainerGroupDiagnosticsLogAnalytics `field:"required" json:"logAnalytics" yaml:"logAnalytics"`
}

