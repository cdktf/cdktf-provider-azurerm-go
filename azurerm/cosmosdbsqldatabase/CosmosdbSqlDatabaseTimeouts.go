package cosmosdbsqldatabase


type CosmosdbSqlDatabaseTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/cosmosdb_sql_database#create CosmosdbSqlDatabase#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/cosmosdb_sql_database#delete CosmosdbSqlDatabase#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/cosmosdb_sql_database#read CosmosdbSqlDatabase#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/cosmosdb_sql_database#update CosmosdbSqlDatabase#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

