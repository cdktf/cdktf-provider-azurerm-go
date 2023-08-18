package cosmosdbsqldatabase


type CosmosdbSqlDatabaseAutoscaleSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/cosmosdb_sql_database#max_throughput CosmosdbSqlDatabase#max_throughput}.
	MaxThroughput *float64 `field:"optional" json:"maxThroughput" yaml:"maxThroughput"`
}

