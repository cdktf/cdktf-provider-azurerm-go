package springcloudservice


type SpringCloudServiceConfigServerGitSettingRepositoryHttpBasicAuth struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.56.0/docs/resources/spring_cloud_service#password SpringCloudService#password}.
	Password *string `field:"required" json:"password" yaml:"password"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.56.0/docs/resources/spring_cloud_service#username SpringCloudService#username}.
	Username *string `field:"required" json:"username" yaml:"username"`
}

