package mssqldatabaseextendedauditingpolicy


type MssqlDatabaseExtendedAuditingPolicyTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_database_extended_auditing_policy#create MssqlDatabaseExtendedAuditingPolicy#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_database_extended_auditing_policy#delete MssqlDatabaseExtendedAuditingPolicy#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_database_extended_auditing_policy#read MssqlDatabaseExtendedAuditingPolicy#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_database_extended_auditing_policy#update MssqlDatabaseExtendedAuditingPolicy#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}
