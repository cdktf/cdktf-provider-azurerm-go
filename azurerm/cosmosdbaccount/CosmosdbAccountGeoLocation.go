package cosmosdbaccount


type CosmosdbAccountGeoLocation struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cosmosdb_account#failover_priority CosmosdbAccount#failover_priority}.
	FailoverPriority *float64 `field:"required" json:"failoverPriority" yaml:"failoverPriority"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cosmosdb_account#location CosmosdbAccount#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cosmosdb_account#zone_redundant CosmosdbAccount#zone_redundant}.
	ZoneRedundant interface{} `field:"optional" json:"zoneRedundant" yaml:"zoneRedundant"`
}

