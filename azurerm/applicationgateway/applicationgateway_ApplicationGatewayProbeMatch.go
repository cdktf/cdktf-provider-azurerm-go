package applicationgateway


type ApplicationGatewayProbeMatch struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/application_gateway#status_code ApplicationGateway#status_code}.
	StatusCode *[]*string `field:"required" json:"statusCode" yaml:"statusCode"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/application_gateway#body ApplicationGateway#body}.
	Body *string `field:"optional" json:"body" yaml:"body"`
}

