package springcloudconfigurationservice


type SpringCloudConfigurationServiceTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/spring_cloud_configuration_service#create SpringCloudConfigurationService#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/spring_cloud_configuration_service#delete SpringCloudConfigurationService#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/spring_cloud_configuration_service#read SpringCloudConfigurationService#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/spring_cloud_configuration_service#update SpringCloudConfigurationService#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

