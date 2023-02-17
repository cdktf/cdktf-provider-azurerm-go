package cosmosdbgremlingraph


type CosmosdbGremlinGraphAutoscaleSettings struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cosmosdb_gremlin_graph#max_throughput CosmosdbGremlinGraph#max_throughput}.
	MaxThroughput *float64 `field:"optional" json:"maxThroughput" yaml:"maxThroughput"`
}

