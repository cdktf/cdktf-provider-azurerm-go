package springcloudgateway


type SpringCloudGatewaySso struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/spring_cloud_gateway#client_id SpringCloudGateway#client_id}.
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/spring_cloud_gateway#client_secret SpringCloudGateway#client_secret}.
	ClientSecret *string `field:"optional" json:"clientSecret" yaml:"clientSecret"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/spring_cloud_gateway#issuer_uri SpringCloudGateway#issuer_uri}.
	IssuerUri *string `field:"optional" json:"issuerUri" yaml:"issuerUri"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/spring_cloud_gateway#scope SpringCloudGateway#scope}.
	Scope *[]*string `field:"optional" json:"scope" yaml:"scope"`
}
