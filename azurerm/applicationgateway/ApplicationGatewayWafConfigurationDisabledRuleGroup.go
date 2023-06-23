package applicationgateway


type ApplicationGatewayWafConfigurationDisabledRuleGroup struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.0/docs/resources/application_gateway#rule_group_name ApplicationGateway#rule_group_name}.
	RuleGroupName *string `field:"required" json:"ruleGroupName" yaml:"ruleGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.0/docs/resources/application_gateway#rules ApplicationGateway#rules}.
	Rules *[]*float64 `field:"optional" json:"rules" yaml:"rules"`
}

