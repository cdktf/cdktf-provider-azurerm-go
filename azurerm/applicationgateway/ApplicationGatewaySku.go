package applicationgateway


type ApplicationGatewaySku struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/application_gateway#name ApplicationGateway#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/application_gateway#tier ApplicationGateway#tier}.
	Tier *string `field:"required" json:"tier" yaml:"tier"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/application_gateway#capacity ApplicationGateway#capacity}.
	Capacity *float64 `field:"optional" json:"capacity" yaml:"capacity"`
}

