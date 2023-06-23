package mysqlflexibleserver


type MysqlFlexibleServerIdentity struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.0/docs/resources/mysql_flexible_server#identity_ids MysqlFlexibleServer#identity_ids}.
	IdentityIds *[]*string `field:"required" json:"identityIds" yaml:"identityIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.0/docs/resources/mysql_flexible_server#type MysqlFlexibleServer#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}

