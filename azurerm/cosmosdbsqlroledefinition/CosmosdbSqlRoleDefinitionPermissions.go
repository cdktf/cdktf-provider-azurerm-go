package cosmosdbsqlroledefinition


type CosmosdbSqlRoleDefinitionPermissions struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cosmosdb_sql_role_definition#data_actions CosmosdbSqlRoleDefinition#data_actions}.
	DataActions *[]*string `field:"required" json:"dataActions" yaml:"dataActions"`
}

