package cosmosdbcassandratable


type CosmosdbCassandraTableSchemaClusterKey struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/cosmosdb_cassandra_table#name CosmosdbCassandraTable#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/cosmosdb_cassandra_table#order_by CosmosdbCassandraTable#order_by}.
	OrderBy *string `field:"required" json:"orderBy" yaml:"orderBy"`
}

