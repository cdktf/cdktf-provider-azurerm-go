package cosmosdbgremlindatabase


type CosmosdbGremlinDatabaseAutoscaleSettings struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cosmosdb_gremlin_database#max_throughput CosmosdbGremlinDatabase#max_throughput}.
	MaxThroughput *float64 `field:"optional" json:"maxThroughput" yaml:"maxThroughput"`
}

