package springcloudgateway


type SpringCloudGatewayQuota struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/spring_cloud_gateway#cpu SpringCloudGateway#cpu}.
	Cpu *string `field:"optional" json:"cpu" yaml:"cpu"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/spring_cloud_gateway#memory SpringCloudGateway#memory}.
	Memory *string `field:"optional" json:"memory" yaml:"memory"`
}

