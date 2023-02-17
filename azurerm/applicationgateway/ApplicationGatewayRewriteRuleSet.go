package applicationgateway


type ApplicationGatewayRewriteRuleSet struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/application_gateway#name ApplicationGateway#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// rewrite_rule block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/application_gateway#rewrite_rule ApplicationGateway#rewrite_rule}
	RewriteRule interface{} `field:"optional" json:"rewriteRule" yaml:"rewriteRule"`
}

