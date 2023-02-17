package cosmosdbsqldatabase


type CosmosdbSqlDatabaseAutoscaleSettings struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cosmosdb_sql_database#max_throughput CosmosdbSqlDatabase#max_throughput}.
	MaxThroughput *float64 `field:"optional" json:"maxThroughput" yaml:"maxThroughput"`
}

