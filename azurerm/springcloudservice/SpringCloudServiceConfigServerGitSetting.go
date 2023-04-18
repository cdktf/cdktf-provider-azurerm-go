package springcloudservice


type SpringCloudServiceConfigServerGitSetting struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.52.0/docs/resources/spring_cloud_service#uri SpringCloudService#uri}.
	Uri *string `field:"required" json:"uri" yaml:"uri"`
	// http_basic_auth block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.52.0/docs/resources/spring_cloud_service#http_basic_auth SpringCloudService#http_basic_auth}
	HttpBasicAuth *SpringCloudServiceConfigServerGitSettingHttpBasicAuth `field:"optional" json:"httpBasicAuth" yaml:"httpBasicAuth"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.52.0/docs/resources/spring_cloud_service#label SpringCloudService#label}.
	Label *string `field:"optional" json:"label" yaml:"label"`
	// repository block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.52.0/docs/resources/spring_cloud_service#repository SpringCloudService#repository}
	Repository interface{} `field:"optional" json:"repository" yaml:"repository"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.52.0/docs/resources/spring_cloud_service#search_paths SpringCloudService#search_paths}.
	SearchPaths *[]*string `field:"optional" json:"searchPaths" yaml:"searchPaths"`
	// ssh_auth block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.52.0/docs/resources/spring_cloud_service#ssh_auth SpringCloudService#ssh_auth}
	SshAuth *SpringCloudServiceConfigServerGitSettingSshAuth `field:"optional" json:"sshAuth" yaml:"sshAuth"`
}

