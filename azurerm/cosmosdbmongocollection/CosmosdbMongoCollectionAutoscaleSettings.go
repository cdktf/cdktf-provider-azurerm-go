package cosmosdbmongocollection


type CosmosdbMongoCollectionAutoscaleSettings struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cosmosdb_mongo_collection#max_throughput CosmosdbMongoCollection#max_throughput}.
	MaxThroughput *float64 `field:"optional" json:"maxThroughput" yaml:"maxThroughput"`
}

