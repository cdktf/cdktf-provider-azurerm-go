package cosmosdbmongocollection


type CosmosdbMongoCollectionIndex struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cosmosdb_mongo_collection#keys CosmosdbMongoCollection#keys}.
	Keys *[]*string `field:"required" json:"keys" yaml:"keys"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cosmosdb_mongo_collection#unique CosmosdbMongoCollection#unique}.
	Unique interface{} `field:"optional" json:"unique" yaml:"unique"`
}

