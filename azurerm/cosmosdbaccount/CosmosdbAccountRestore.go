package cosmosdbaccount


type CosmosdbAccountRestore struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.56.0/docs/resources/cosmosdb_account#restore_timestamp_in_utc CosmosdbAccount#restore_timestamp_in_utc}.
	RestoreTimestampInUtc *string `field:"required" json:"restoreTimestampInUtc" yaml:"restoreTimestampInUtc"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.56.0/docs/resources/cosmosdb_account#source_cosmosdb_account_id CosmosdbAccount#source_cosmosdb_account_id}.
	SourceCosmosdbAccountId *string `field:"required" json:"sourceCosmosdbAccountId" yaml:"sourceCosmosdbAccountId"`
	// database block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.56.0/docs/resources/cosmosdb_account#database CosmosdbAccount#database}
	Database interface{} `field:"optional" json:"database" yaml:"database"`
}

