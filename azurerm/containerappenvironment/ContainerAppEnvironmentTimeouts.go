package containerappenvironment


type ContainerAppEnvironmentTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_app_environment#create ContainerAppEnvironment#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_app_environment#delete ContainerAppEnvironment#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_app_environment#read ContainerAppEnvironment#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_app_environment#update ContainerAppEnvironment#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}
