package postgresqlflexibleserver


type PostgresqlFlexibleServerTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/postgresql_flexible_server#create PostgresqlFlexibleServer#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/postgresql_flexible_server#delete PostgresqlFlexibleServer#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/postgresql_flexible_server#read PostgresqlFlexibleServer#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/postgresql_flexible_server#update PostgresqlFlexibleServer#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

