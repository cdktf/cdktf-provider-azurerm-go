package cosmosdbcassandratable


type CosmosdbCassandraTableSchemaColumn struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cosmosdb_cassandra_table#name CosmosdbCassandraTable#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cosmosdb_cassandra_table#type CosmosdbCassandraTable#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}

