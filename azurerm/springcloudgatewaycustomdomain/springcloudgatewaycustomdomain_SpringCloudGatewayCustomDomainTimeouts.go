package springcloudgatewaycustomdomain


type SpringCloudGatewayCustomDomainTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/spring_cloud_gateway_custom_domain#create SpringCloudGatewayCustomDomain#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/spring_cloud_gateway_custom_domain#delete SpringCloudGatewayCustomDomain#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/spring_cloud_gateway_custom_domain#read SpringCloudGatewayCustomDomain#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/spring_cloud_gateway_custom_domain#update SpringCloudGatewayCustomDomain#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}
