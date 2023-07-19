package cosmosdbpostgresqlnodeconfiguration


type CosmosdbPostgresqlNodeConfigurationTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/cosmosdb_postgresql_node_configuration#create CosmosdbPostgresqlNodeConfiguration#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/cosmosdb_postgresql_node_configuration#delete CosmosdbPostgresqlNodeConfiguration#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/cosmosdb_postgresql_node_configuration#read CosmosdbPostgresqlNodeConfiguration#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/cosmosdb_postgresql_node_configuration#update CosmosdbPostgresqlNodeConfiguration#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

