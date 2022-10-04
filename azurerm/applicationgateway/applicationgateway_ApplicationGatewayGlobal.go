package applicationgateway


type ApplicationGatewayGlobal struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/application_gateway#request_buffering_enabled ApplicationGateway#request_buffering_enabled}.
	RequestBufferingEnabled interface{} `field:"required" json:"requestBufferingEnabled" yaml:"requestBufferingEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/application_gateway#response_buffering_enabled ApplicationGateway#response_buffering_enabled}.
	ResponseBufferingEnabled interface{} `field:"required" json:"responseBufferingEnabled" yaml:"responseBufferingEnabled"`
}

