package cosmosdbsqldedicatedgateway


type CosmosdbSqlDedicatedGatewayTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cosmosdb_sql_dedicated_gateway#create CosmosdbSqlDedicatedGateway#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cosmosdb_sql_dedicated_gateway#delete CosmosdbSqlDedicatedGateway#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cosmosdb_sql_dedicated_gateway#read CosmosdbSqlDedicatedGateway#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cosmosdb_sql_dedicated_gateway#update CosmosdbSqlDedicatedGateway#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}
