package cosmosdbaccount


type CosmosdbAccountConsistencyPolicy struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cosmosdb_account#consistency_level CosmosdbAccount#consistency_level}.
	ConsistencyLevel *string `field:"required" json:"consistencyLevel" yaml:"consistencyLevel"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cosmosdb_account#max_interval_in_seconds CosmosdbAccount#max_interval_in_seconds}.
	MaxIntervalInSeconds *float64 `field:"optional" json:"maxIntervalInSeconds" yaml:"maxIntervalInSeconds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cosmosdb_account#max_staleness_prefix CosmosdbAccount#max_staleness_prefix}.
	MaxStalenessPrefix *float64 `field:"optional" json:"maxStalenessPrefix" yaml:"maxStalenessPrefix"`
}

