package springcloudgatewayrouteconfig


type SpringCloudGatewayRouteConfigTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/spring_cloud_gateway_route_config#create SpringCloudGatewayRouteConfig#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/spring_cloud_gateway_route_config#delete SpringCloudGatewayRouteConfig#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/spring_cloud_gateway_route_config#read SpringCloudGatewayRouteConfig#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/spring_cloud_gateway_route_config#update SpringCloudGatewayRouteConfig#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}
