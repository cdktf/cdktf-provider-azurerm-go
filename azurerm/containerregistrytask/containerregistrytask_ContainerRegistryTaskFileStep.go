package containerregistrytask


type ContainerRegistryTaskFileStep struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_registry_task#task_file_path ContainerRegistryTask#task_file_path}.
	TaskFilePath *string `field:"required" json:"taskFilePath" yaml:"taskFilePath"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_registry_task#context_access_token ContainerRegistryTask#context_access_token}.
	ContextAccessToken *string `field:"optional" json:"contextAccessToken" yaml:"contextAccessToken"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_registry_task#context_path ContainerRegistryTask#context_path}.
	ContextPath *string `field:"optional" json:"contextPath" yaml:"contextPath"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_registry_task#secret_values ContainerRegistryTask#secret_values}.
	SecretValues *map[string]*string `field:"optional" json:"secretValues" yaml:"secretValues"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_registry_task#value_file_path ContainerRegistryTask#value_file_path}.
	ValueFilePath *string `field:"optional" json:"valueFilePath" yaml:"valueFilePath"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_registry_task#values ContainerRegistryTask#values}.
	Values *map[string]*string `field:"optional" json:"values" yaml:"values"`
}

