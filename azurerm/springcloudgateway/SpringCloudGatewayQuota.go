package springcloudgateway


type SpringCloudGatewayQuota struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.68.0/docs/resources/spring_cloud_gateway#cpu SpringCloudGateway#cpu}.
	Cpu *string `field:"optional" json:"cpu" yaml:"cpu"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.68.0/docs/resources/spring_cloud_gateway#memory SpringCloudGateway#memory}.
	Memory *string `field:"optional" json:"memory" yaml:"memory"`
}

