package servicebusnamespacenetworkruleset


type ServicebusNamespaceNetworkRuleSetTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/servicebus_namespace_network_rule_set#create ServicebusNamespaceNetworkRuleSet#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/servicebus_namespace_network_rule_set#delete ServicebusNamespaceNetworkRuleSet#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/servicebus_namespace_network_rule_set#read ServicebusNamespaceNetworkRuleSet#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/servicebus_namespace_network_rule_set#update ServicebusNamespaceNetworkRuleSet#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

