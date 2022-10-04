package cosmosdbsqlcontainer


type CosmosdbSqlContainerUniqueKey struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cosmosdb_sql_container#paths CosmosdbSqlContainer#paths}.
	Paths *[]*string `field:"required" json:"paths" yaml:"paths"`
}

