package applicationgateway


type ApplicationGatewayRewriteRuleSetRewriteRuleResponseHeaderConfiguration struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/application_gateway#header_name ApplicationGateway#header_name}.
	HeaderName *string `field:"required" json:"headerName" yaml:"headerName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/application_gateway#header_value ApplicationGateway#header_value}.
	HeaderValue *string `field:"required" json:"headerValue" yaml:"headerValue"`
}

