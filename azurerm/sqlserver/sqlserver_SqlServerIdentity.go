package sqlserver


type SqlServerIdentity struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/sql_server#type SqlServer#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}

