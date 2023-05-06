package containerapp


type ContainerAppRegistry struct {
	// The hostname for the Container Registry.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/container_app#server ContainerApp#server}
	Server *string `field:"required" json:"server" yaml:"server"`
	// ID of the System or User Managed Identity used to pull images from the Container Registry.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/container_app#identity ContainerApp#identity}
	Identity *string `field:"optional" json:"identity" yaml:"identity"`
	// The name of the Secret Reference containing the password value for this user on the Container Registry.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/container_app#password_secret_name ContainerApp#password_secret_name}
	PasswordSecretName *string `field:"optional" json:"passwordSecretName" yaml:"passwordSecretName"`
	// The username to use for this Container Registry.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/container_app#username ContainerApp#username}
	Username *string `field:"optional" json:"username" yaml:"username"`
}

