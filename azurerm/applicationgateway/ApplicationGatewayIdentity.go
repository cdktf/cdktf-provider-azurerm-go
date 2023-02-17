package applicationgateway


type ApplicationGatewayIdentity struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/application_gateway#identity_ids ApplicationGateway#identity_ids}.
	IdentityIds *[]*string `field:"required" json:"identityIds" yaml:"identityIds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/application_gateway#type ApplicationGateway#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}

