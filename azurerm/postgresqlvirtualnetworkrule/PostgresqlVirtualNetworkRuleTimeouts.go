package postgresqlvirtualnetworkrule


type PostgresqlVirtualNetworkRuleTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/postgresql_virtual_network_rule#create PostgresqlVirtualNetworkRule#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/postgresql_virtual_network_rule#delete PostgresqlVirtualNetworkRule#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/postgresql_virtual_network_rule#read PostgresqlVirtualNetworkRule#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/postgresql_virtual_network_rule#update PostgresqlVirtualNetworkRule#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

