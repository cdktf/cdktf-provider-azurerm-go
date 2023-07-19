package cosmosdbmongodatabase


type CosmosdbMongoDatabaseAutoscaleSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/cosmosdb_mongo_database#max_throughput CosmosdbMongoDatabase#max_throughput}.
	MaxThroughput *float64 `field:"optional" json:"maxThroughput" yaml:"maxThroughput"`
}

