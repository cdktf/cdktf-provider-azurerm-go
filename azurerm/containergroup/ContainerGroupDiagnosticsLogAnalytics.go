package containergroup


type ContainerGroupDiagnosticsLogAnalytics struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_group#workspace_id ContainerGroup#workspace_id}.
	WorkspaceId *string `field:"required" json:"workspaceId" yaml:"workspaceId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_group#workspace_key ContainerGroup#workspace_key}.
	WorkspaceKey *string `field:"required" json:"workspaceKey" yaml:"workspaceKey"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_group#log_type ContainerGroup#log_type}.
	LogType *string `field:"optional" json:"logType" yaml:"logType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_group#metadata ContainerGroup#metadata}.
	Metadata *map[string]*string `field:"optional" json:"metadata" yaml:"metadata"`
}
