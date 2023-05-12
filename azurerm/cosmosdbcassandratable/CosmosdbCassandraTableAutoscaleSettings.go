package cosmosdbcassandratable


type CosmosdbCassandraTableAutoscaleSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.56.0/docs/resources/cosmosdb_cassandra_table#max_throughput CosmosdbCassandraTable#max_throughput}.
	MaxThroughput *float64 `field:"optional" json:"maxThroughput" yaml:"maxThroughput"`
}

