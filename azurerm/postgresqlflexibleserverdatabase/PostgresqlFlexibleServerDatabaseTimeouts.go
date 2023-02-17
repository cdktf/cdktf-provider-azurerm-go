package postgresqlflexibleserverdatabase


type PostgresqlFlexibleServerDatabaseTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/postgresql_flexible_server_database#create PostgresqlFlexibleServerDatabase#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/postgresql_flexible_server_database#delete PostgresqlFlexibleServerDatabase#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/postgresql_flexible_server_database#read PostgresqlFlexibleServerDatabase#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}

