package applicationgateway


type ApplicationGatewayWafConfigurationExclusion struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/application_gateway#match_variable ApplicationGateway#match_variable}.
	MatchVariable *string `field:"required" json:"matchVariable" yaml:"matchVariable"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/application_gateway#selector ApplicationGateway#selector}.
	Selector *string `field:"optional" json:"selector" yaml:"selector"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/application_gateway#selector_match_operator ApplicationGateway#selector_match_operator}.
	SelectorMatchOperator *string `field:"optional" json:"selectorMatchOperator" yaml:"selectorMatchOperator"`
}

