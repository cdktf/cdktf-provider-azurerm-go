package cosmosdbaccount


type CosmosdbAccountRestoreDatabase struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cosmosdb_account#name CosmosdbAccount#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cosmosdb_account#collection_names CosmosdbAccount#collection_names}.
	CollectionNames *[]*string `field:"optional" json:"collectionNames" yaml:"collectionNames"`
}

