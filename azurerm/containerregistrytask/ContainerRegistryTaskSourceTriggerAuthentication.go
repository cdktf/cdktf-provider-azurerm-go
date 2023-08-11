package containerregistrytask


type ContainerRegistryTaskSourceTriggerAuthentication struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.69.0/docs/resources/container_registry_task#token ContainerRegistryTask#token}.
	Token *string `field:"required" json:"token" yaml:"token"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.69.0/docs/resources/container_registry_task#token_type ContainerRegistryTask#token_type}.
	TokenType *string `field:"required" json:"tokenType" yaml:"tokenType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.69.0/docs/resources/container_registry_task#expire_in_seconds ContainerRegistryTask#expire_in_seconds}.
	ExpireInSeconds *float64 `field:"optional" json:"expireInSeconds" yaml:"expireInSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.69.0/docs/resources/container_registry_task#refresh_token ContainerRegistryTask#refresh_token}.
	RefreshToken *string `field:"optional" json:"refreshToken" yaml:"refreshToken"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.69.0/docs/resources/container_registry_task#scope ContainerRegistryTask#scope}.
	Scope *string `field:"optional" json:"scope" yaml:"scope"`
}

