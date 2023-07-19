package springcloudconfigurationservice


type SpringCloudConfigurationServiceRepository struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/spring_cloud_configuration_service#label SpringCloudConfigurationService#label}.
	Label *string `field:"required" json:"label" yaml:"label"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/spring_cloud_configuration_service#name SpringCloudConfigurationService#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/spring_cloud_configuration_service#patterns SpringCloudConfigurationService#patterns}.
	Patterns *[]*string `field:"required" json:"patterns" yaml:"patterns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/spring_cloud_configuration_service#uri SpringCloudConfigurationService#uri}.
	Uri *string `field:"required" json:"uri" yaml:"uri"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/spring_cloud_configuration_service#host_key SpringCloudConfigurationService#host_key}.
	HostKey *string `field:"optional" json:"hostKey" yaml:"hostKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/spring_cloud_configuration_service#host_key_algorithm SpringCloudConfigurationService#host_key_algorithm}.
	HostKeyAlgorithm *string `field:"optional" json:"hostKeyAlgorithm" yaml:"hostKeyAlgorithm"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/spring_cloud_configuration_service#password SpringCloudConfigurationService#password}.
	Password *string `field:"optional" json:"password" yaml:"password"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/spring_cloud_configuration_service#private_key SpringCloudConfigurationService#private_key}.
	PrivateKey *string `field:"optional" json:"privateKey" yaml:"privateKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/spring_cloud_configuration_service#search_paths SpringCloudConfigurationService#search_paths}.
	SearchPaths *[]*string `field:"optional" json:"searchPaths" yaml:"searchPaths"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/spring_cloud_configuration_service#strict_host_key_checking SpringCloudConfigurationService#strict_host_key_checking}.
	StrictHostKeyChecking interface{} `field:"optional" json:"strictHostKeyChecking" yaml:"strictHostKeyChecking"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/spring_cloud_configuration_service#username SpringCloudConfigurationService#username}.
	Username *string `field:"optional" json:"username" yaml:"username"`
}

