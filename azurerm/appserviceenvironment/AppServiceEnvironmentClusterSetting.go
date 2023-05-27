package appserviceenvironment


type AppServiceEnvironmentClusterSetting struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.58.0/docs/resources/app_service_environment#name AppServiceEnvironment#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.58.0/docs/resources/app_service_environment#value AppServiceEnvironment#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

