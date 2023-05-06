package cosmosdbtable


type CosmosdbTableAutoscaleSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/cosmosdb_table#max_throughput CosmosdbTable#max_throughput}.
	MaxThroughput *float64 `field:"optional" json:"maxThroughput" yaml:"maxThroughput"`
}

