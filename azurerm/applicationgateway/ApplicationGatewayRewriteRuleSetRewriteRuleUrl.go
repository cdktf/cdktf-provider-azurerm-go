package applicationgateway


type ApplicationGatewayRewriteRuleSetRewriteRuleUrl struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/application_gateway#components ApplicationGateway#components}.
	Components *string `field:"optional" json:"components" yaml:"components"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/application_gateway#path ApplicationGateway#path}.
	Path *string `field:"optional" json:"path" yaml:"path"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/application_gateway#query_string ApplicationGateway#query_string}.
	QueryString *string `field:"optional" json:"queryString" yaml:"queryString"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/application_gateway#reroute ApplicationGateway#reroute}.
	Reroute interface{} `field:"optional" json:"reroute" yaml:"reroute"`
}

