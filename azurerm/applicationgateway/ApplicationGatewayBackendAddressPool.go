package applicationgateway


type ApplicationGatewayBackendAddressPool struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/application_gateway#name ApplicationGateway#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/application_gateway#fqdns ApplicationGateway#fqdns}.
	Fqdns *[]*string `field:"optional" json:"fqdns" yaml:"fqdns"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/application_gateway#ip_addresses ApplicationGateway#ip_addresses}.
	IpAddresses *[]*string `field:"optional" json:"ipAddresses" yaml:"ipAddresses"`
}

