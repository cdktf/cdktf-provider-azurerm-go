package postgresqlflexibleserver


type PostgresqlFlexibleServerIdentity struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.61.0/docs/resources/postgresql_flexible_server#identity_ids PostgresqlFlexibleServer#identity_ids}.
	IdentityIds *[]*string `field:"required" json:"identityIds" yaml:"identityIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.61.0/docs/resources/postgresql_flexible_server#type PostgresqlFlexibleServer#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}

