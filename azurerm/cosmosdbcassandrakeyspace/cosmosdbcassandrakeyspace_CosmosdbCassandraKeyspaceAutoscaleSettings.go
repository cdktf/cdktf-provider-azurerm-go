package cosmosdbcassandrakeyspace


type CosmosdbCassandraKeyspaceAutoscaleSettings struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cosmosdb_cassandra_keyspace#max_throughput CosmosdbCassandraKeyspace#max_throughput}.
	MaxThroughput *float64 `field:"optional" json:"maxThroughput" yaml:"maxThroughput"`
}
