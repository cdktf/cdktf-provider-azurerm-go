// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cosmosdbsqlcontainer


type CosmosdbSqlContainerUniqueKey struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/cosmosdb_sql_container#paths CosmosdbSqlContainer#paths}.
	Paths *[]*string `field:"required" json:"paths" yaml:"paths"`
}

