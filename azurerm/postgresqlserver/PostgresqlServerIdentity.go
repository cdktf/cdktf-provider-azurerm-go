package postgresqlserver


type PostgresqlServerIdentity struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/postgresql_server#type PostgresqlServer#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}

