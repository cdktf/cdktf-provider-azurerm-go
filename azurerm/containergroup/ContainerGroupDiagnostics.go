package containergroup


type ContainerGroupDiagnostics struct {
	// log_analytics block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_group#log_analytics ContainerGroup#log_analytics}
	LogAnalytics *ContainerGroupDiagnosticsLogAnalytics `field:"required" json:"logAnalytics" yaml:"logAnalytics"`
}
