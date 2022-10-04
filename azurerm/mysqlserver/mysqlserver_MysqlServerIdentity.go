package mysqlserver


type MysqlServerIdentity struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mysql_server#type MysqlServer#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}

