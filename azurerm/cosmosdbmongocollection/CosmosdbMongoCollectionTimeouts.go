package cosmosdbmongocollection


type CosmosdbMongoCollectionTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cosmosdb_mongo_collection#create CosmosdbMongoCollection#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cosmosdb_mongo_collection#delete CosmosdbMongoCollection#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cosmosdb_mongo_collection#read CosmosdbMongoCollection#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cosmosdb_mongo_collection#update CosmosdbMongoCollection#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}
