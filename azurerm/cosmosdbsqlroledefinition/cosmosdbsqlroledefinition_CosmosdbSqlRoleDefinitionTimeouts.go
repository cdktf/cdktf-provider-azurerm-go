package cosmosdbsqlroledefinition


type CosmosdbSqlRoleDefinitionTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cosmosdb_sql_role_definition#create CosmosdbSqlRoleDefinition#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cosmosdb_sql_role_definition#delete CosmosdbSqlRoleDefinition#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cosmosdb_sql_role_definition#read CosmosdbSqlRoleDefinition#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cosmosdb_sql_role_definition#update CosmosdbSqlRoleDefinition#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}
