package applicationgateway


type ApplicationGatewayRewriteRuleSet struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/application_gateway#name ApplicationGateway#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// rewrite_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/application_gateway#rewrite_rule ApplicationGateway#rewrite_rule}
	RewriteRule interface{} `field:"optional" json:"rewriteRule" yaml:"rewriteRule"`
}

