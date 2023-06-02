package firewallnetworkrulecollection


type FirewallNetworkRuleCollectionTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.59.0/docs/resources/firewall_network_rule_collection#create FirewallNetworkRuleCollection#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.59.0/docs/resources/firewall_network_rule_collection#delete FirewallNetworkRuleCollection#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.59.0/docs/resources/firewall_network_rule_collection#read FirewallNetworkRuleCollection#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.59.0/docs/resources/firewall_network_rule_collection#update FirewallNetworkRuleCollection#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

