package postgresqlflexibleserverfirewallrule


type PostgresqlFlexibleServerFirewallRuleTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/postgresql_flexible_server_firewall_rule#create PostgresqlFlexibleServerFirewallRule#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/postgresql_flexible_server_firewall_rule#delete PostgresqlFlexibleServerFirewallRule#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/postgresql_flexible_server_firewall_rule#read PostgresqlFlexibleServerFirewallRule#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/postgresql_flexible_server_firewall_rule#update PostgresqlFlexibleServerFirewallRule#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}
