// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cosmosdbaccount


type CosmosdbAccountBackup struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.12.0/docs/resources/cosmosdb_account#type CosmosdbAccount#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.12.0/docs/resources/cosmosdb_account#interval_in_minutes CosmosdbAccount#interval_in_minutes}.
	IntervalInMinutes *float64 `field:"optional" json:"intervalInMinutes" yaml:"intervalInMinutes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.12.0/docs/resources/cosmosdb_account#retention_in_hours CosmosdbAccount#retention_in_hours}.
	RetentionInHours *float64 `field:"optional" json:"retentionInHours" yaml:"retentionInHours"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.12.0/docs/resources/cosmosdb_account#storage_redundancy CosmosdbAccount#storage_redundancy}.
	StorageRedundancy *string `field:"optional" json:"storageRedundancy" yaml:"storageRedundancy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.12.0/docs/resources/cosmosdb_account#tier CosmosdbAccount#tier}.
	Tier *string `field:"optional" json:"tier" yaml:"tier"`
}

