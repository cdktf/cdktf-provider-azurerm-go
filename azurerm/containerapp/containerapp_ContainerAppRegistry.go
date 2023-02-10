package containerapp


type ContainerAppRegistry struct {
	// The name of the Secret Reference containing the password value for this user on the Container Registry.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_app#password_secret_name ContainerApp#password_secret_name}
	PasswordSecretName *string `field:"required" json:"passwordSecretName" yaml:"passwordSecretName"`
	// The hostname for the Container Registry.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_app#server ContainerApp#server}
	Server *string `field:"required" json:"server" yaml:"server"`
	// The username to use for this Container Registry.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_app#username ContainerApp#username}
	Username *string `field:"required" json:"username" yaml:"username"`
}

