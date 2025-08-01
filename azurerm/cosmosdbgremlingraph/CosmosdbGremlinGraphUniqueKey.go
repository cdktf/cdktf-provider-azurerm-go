// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cosmosdbgremlingraph


type CosmosdbGremlinGraphUniqueKey struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/cosmosdb_gremlin_graph#paths CosmosdbGremlinGraph#paths}.
	Paths *[]*string `field:"required" json:"paths" yaml:"paths"`
}

