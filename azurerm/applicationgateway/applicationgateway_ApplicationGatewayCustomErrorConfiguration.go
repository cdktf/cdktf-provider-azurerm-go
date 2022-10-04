package applicationgateway


type ApplicationGatewayCustomErrorConfiguration struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/application_gateway#custom_error_page_url ApplicationGateway#custom_error_page_url}.
	CustomErrorPageUrl *string `field:"required" json:"customErrorPageUrl" yaml:"customErrorPageUrl"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/application_gateway#status_code ApplicationGateway#status_code}.
	StatusCode *string `field:"required" json:"statusCode" yaml:"statusCode"`
}

