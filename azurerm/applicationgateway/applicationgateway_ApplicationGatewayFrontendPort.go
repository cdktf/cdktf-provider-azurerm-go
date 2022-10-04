package applicationgateway


type ApplicationGatewayFrontendPort struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/application_gateway#name ApplicationGateway#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/application_gateway#port ApplicationGateway#port}.
	Port *float64 `field:"required" json:"port" yaml:"port"`
}

