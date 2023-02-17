package firewallapplicationrulecollection


type FirewallApplicationRuleCollectionRuleProtocol struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/firewall_application_rule_collection#port FirewallApplicationRuleCollection#port}.
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/firewall_application_rule_collection#type FirewallApplicationRuleCollection#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}

