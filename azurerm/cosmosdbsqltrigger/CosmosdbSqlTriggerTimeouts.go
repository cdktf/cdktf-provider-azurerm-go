package cosmosdbsqltrigger


type CosmosdbSqlTriggerTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/cosmosdb_sql_trigger#create CosmosdbSqlTrigger#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/cosmosdb_sql_trigger#delete CosmosdbSqlTrigger#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/cosmosdb_sql_trigger#read CosmosdbSqlTrigger#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/cosmosdb_sql_trigger#update CosmosdbSqlTrigger#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

