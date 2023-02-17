package springcloudbuilddeployment


type SpringCloudBuildDeploymentQuota struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/spring_cloud_build_deployment#cpu SpringCloudBuildDeployment#cpu}.
	Cpu *string `field:"optional" json:"cpu" yaml:"cpu"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/spring_cloud_build_deployment#memory SpringCloudBuildDeployment#memory}.
	Memory *string `field:"optional" json:"memory" yaml:"memory"`
}

