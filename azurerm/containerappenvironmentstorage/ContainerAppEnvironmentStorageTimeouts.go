package containerappenvironmentstorage


type ContainerAppEnvironmentStorageTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_app_environment_storage#create ContainerAppEnvironmentStorage#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_app_environment_storage#delete ContainerAppEnvironmentStorage#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_app_environment_storage#read ContainerAppEnvironmentStorage#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_app_environment_storage#update ContainerAppEnvironmentStorage#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}
