package cosmosdbsqlcontainer


type CosmosdbSqlContainerTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/cosmosdb_sql_container#create CosmosdbSqlContainer#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/cosmosdb_sql_container#delete CosmosdbSqlContainer#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/cosmosdb_sql_container#read CosmosdbSqlContainer#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/cosmosdb_sql_container#update CosmosdbSqlContainer#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

