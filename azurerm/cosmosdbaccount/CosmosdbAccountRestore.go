// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cosmosdbaccount


type CosmosdbAccountRestore struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.49.0/docs/resources/cosmosdb_account#restore_timestamp_in_utc CosmosdbAccount#restore_timestamp_in_utc}.
	RestoreTimestampInUtc *string `field:"required" json:"restoreTimestampInUtc" yaml:"restoreTimestampInUtc"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.49.0/docs/resources/cosmosdb_account#source_cosmosdb_account_id CosmosdbAccount#source_cosmosdb_account_id}.
	SourceCosmosdbAccountId *string `field:"required" json:"sourceCosmosdbAccountId" yaml:"sourceCosmosdbAccountId"`
	// database block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.49.0/docs/resources/cosmosdb_account#database CosmosdbAccount#database}
	Database interface{} `field:"optional" json:"database" yaml:"database"`
	// gremlin_database block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.49.0/docs/resources/cosmosdb_account#gremlin_database CosmosdbAccount#gremlin_database}
	GremlinDatabase interface{} `field:"optional" json:"gremlinDatabase" yaml:"gremlinDatabase"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.49.0/docs/resources/cosmosdb_account#tables_to_restore CosmosdbAccount#tables_to_restore}.
	TablesToRestore *[]*string `field:"optional" json:"tablesToRestore" yaml:"tablesToRestore"`
}

