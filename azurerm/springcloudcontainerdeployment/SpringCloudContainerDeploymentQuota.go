package springcloudcontainerdeployment


type SpringCloudContainerDeploymentQuota struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/spring_cloud_container_deployment#cpu SpringCloudContainerDeployment#cpu}.
	Cpu *string `field:"optional" json:"cpu" yaml:"cpu"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/spring_cloud_container_deployment#memory SpringCloudContainerDeployment#memory}.
	Memory *string `field:"optional" json:"memory" yaml:"memory"`
}

