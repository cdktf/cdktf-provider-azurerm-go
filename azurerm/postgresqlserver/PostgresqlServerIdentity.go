package postgresqlserver


type PostgresqlServerIdentity struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.69.0/docs/resources/postgresql_server#type PostgresqlServer#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}

