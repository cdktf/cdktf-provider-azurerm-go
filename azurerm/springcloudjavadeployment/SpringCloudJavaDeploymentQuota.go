package springcloudjavadeployment


type SpringCloudJavaDeploymentQuota struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/spring_cloud_java_deployment#cpu SpringCloudJavaDeployment#cpu}.
	Cpu *string `field:"optional" json:"cpu" yaml:"cpu"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/spring_cloud_java_deployment#memory SpringCloudJavaDeployment#memory}.
	Memory *string `field:"optional" json:"memory" yaml:"memory"`
}

