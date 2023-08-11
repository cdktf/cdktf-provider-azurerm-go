package applicationgateway


type ApplicationGatewayGlobal struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.69.0/docs/resources/application_gateway#request_buffering_enabled ApplicationGateway#request_buffering_enabled}.
	RequestBufferingEnabled interface{} `field:"required" json:"requestBufferingEnabled" yaml:"requestBufferingEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.69.0/docs/resources/application_gateway#response_buffering_enabled ApplicationGateway#response_buffering_enabled}.
	ResponseBufferingEnabled interface{} `field:"required" json:"responseBufferingEnabled" yaml:"responseBufferingEnabled"`
}

