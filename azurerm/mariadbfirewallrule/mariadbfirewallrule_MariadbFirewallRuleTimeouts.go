package mariadbfirewallrule


type MariadbFirewallRuleTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mariadb_firewall_rule#create MariadbFirewallRule#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mariadb_firewall_rule#delete MariadbFirewallRule#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mariadb_firewall_rule#read MariadbFirewallRule#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mariadb_firewall_rule#update MariadbFirewallRule#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

