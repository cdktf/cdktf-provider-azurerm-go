package applicationgateway


type ApplicationGatewayRewriteRuleSetRewriteRuleCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/application_gateway#pattern ApplicationGateway#pattern}.
	Pattern *string `field:"required" json:"pattern" yaml:"pattern"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/application_gateway#variable ApplicationGateway#variable}.
	Variable *string `field:"required" json:"variable" yaml:"variable"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/application_gateway#ignore_case ApplicationGateway#ignore_case}.
	IgnoreCase interface{} `field:"optional" json:"ignoreCase" yaml:"ignoreCase"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/application_gateway#negate ApplicationGateway#negate}.
	Negate interface{} `field:"optional" json:"negate" yaml:"negate"`
}

