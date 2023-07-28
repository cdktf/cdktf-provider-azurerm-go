package appservicesourcecontrolslot


type AppServiceSourceControlSlotGithubActionConfigurationContainerConfiguration struct {
	// The image name for the build.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/app_service_source_control_slot#image_name AppServiceSourceControlSlot#image_name}
	ImageName *string `field:"required" json:"imageName" yaml:"imageName"`
	// The server URL for the container registry where the build will be hosted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/app_service_source_control_slot#registry_url AppServiceSourceControlSlot#registry_url}
	RegistryUrl *string `field:"required" json:"registryUrl" yaml:"registryUrl"`
	// The password used to upload the image to the container registry.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/app_service_source_control_slot#registry_password AppServiceSourceControlSlot#registry_password}
	RegistryPassword *string `field:"optional" json:"registryPassword" yaml:"registryPassword"`
	// The username used to upload the image to the container registry.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/app_service_source_control_slot#registry_username AppServiceSourceControlSlot#registry_username}
	RegistryUsername *string `field:"optional" json:"registryUsername" yaml:"registryUsername"`
}

