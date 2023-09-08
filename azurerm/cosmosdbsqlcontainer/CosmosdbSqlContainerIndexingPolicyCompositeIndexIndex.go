// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cosmosdbsqlcontainer


type CosmosdbSqlContainerIndexingPolicyCompositeIndexIndex struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.72.0/docs/resources/cosmosdb_sql_container#order CosmosdbSqlContainer#order}.
	Order *string `field:"required" json:"order" yaml:"order"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.72.0/docs/resources/cosmosdb_sql_container#path CosmosdbSqlContainer#path}.
	Path *string `field:"required" json:"path" yaml:"path"`
}

