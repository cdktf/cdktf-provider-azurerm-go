package cdnendpoint


type CdnEndpointDeliveryRuleUrlFileNameCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#operator CdnEndpoint#operator}.
	Operator *string `field:"required" json:"operator" yaml:"operator"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#match_values CdnEndpoint#match_values}.
	MatchValues *[]*string `field:"optional" json:"matchValues" yaml:"matchValues"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#negate_condition CdnEndpoint#negate_condition}.
	NegateCondition interface{} `field:"optional" json:"negateCondition" yaml:"negateCondition"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_endpoint#transforms CdnEndpoint#transforms}.
	Transforms *[]*string `field:"optional" json:"transforms" yaml:"transforms"`
}
