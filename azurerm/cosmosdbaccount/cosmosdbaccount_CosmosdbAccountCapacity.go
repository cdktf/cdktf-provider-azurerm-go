package cosmosdbaccount


type CosmosdbAccountCapacity struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cosmosdb_account#total_throughput_limit CosmosdbAccount#total_throughput_limit}.
	TotalThroughputLimit *float64 `field:"required" json:"totalThroughputLimit" yaml:"totalThroughputLimit"`
}

