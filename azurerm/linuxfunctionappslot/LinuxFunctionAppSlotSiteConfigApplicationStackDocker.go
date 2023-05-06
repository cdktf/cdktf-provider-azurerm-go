package linuxfunctionappslot


type LinuxFunctionAppSlotSiteConfigApplicationStackDocker struct {
	// The name of the Docker image to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/linux_function_app_slot#image_name LinuxFunctionAppSlot#image_name}
	ImageName *string `field:"required" json:"imageName" yaml:"imageName"`
	// The image tag of the image to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/linux_function_app_slot#image_tag LinuxFunctionAppSlot#image_tag}
	ImageTag *string `field:"required" json:"imageTag" yaml:"imageTag"`
	// The URL of the docker registry.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/linux_function_app_slot#registry_url LinuxFunctionAppSlot#registry_url}
	RegistryUrl *string `field:"required" json:"registryUrl" yaml:"registryUrl"`
	// The password for the account to use to connect to the registry.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/linux_function_app_slot#registry_password LinuxFunctionAppSlot#registry_password}
	RegistryPassword *string `field:"optional" json:"registryPassword" yaml:"registryPassword"`
	// The username to use for connections to the registry.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/linux_function_app_slot#registry_username LinuxFunctionAppSlot#registry_username}
	RegistryUsername *string `field:"optional" json:"registryUsername" yaml:"registryUsername"`
}

