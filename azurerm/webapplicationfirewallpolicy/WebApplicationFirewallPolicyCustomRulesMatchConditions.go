package webapplicationfirewallpolicy


type WebApplicationFirewallPolicyCustomRulesMatchConditions struct {
	// match_variables block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/web_application_firewall_policy#match_variables WebApplicationFirewallPolicy#match_variables}
	MatchVariables interface{} `field:"required" json:"matchVariables" yaml:"matchVariables"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/web_application_firewall_policy#operator WebApplicationFirewallPolicy#operator}.
	Operator *string `field:"required" json:"operator" yaml:"operator"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/web_application_firewall_policy#match_values WebApplicationFirewallPolicy#match_values}.
	MatchValues *[]*string `field:"optional" json:"matchValues" yaml:"matchValues"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/web_application_firewall_policy#negation_condition WebApplicationFirewallPolicy#negation_condition}.
	NegationCondition interface{} `field:"optional" json:"negationCondition" yaml:"negationCondition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/web_application_firewall_policy#transforms WebApplicationFirewallPolicy#transforms}.
	Transforms *[]*string `field:"optional" json:"transforms" yaml:"transforms"`
}

