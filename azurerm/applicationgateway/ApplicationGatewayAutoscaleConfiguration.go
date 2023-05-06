package applicationgateway


type ApplicationGatewayAutoscaleConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/application_gateway#min_capacity ApplicationGateway#min_capacity}.
	MinCapacity *float64 `field:"required" json:"minCapacity" yaml:"minCapacity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/application_gateway#max_capacity ApplicationGateway#max_capacity}.
	MaxCapacity *float64 `field:"optional" json:"maxCapacity" yaml:"maxCapacity"`
}

