// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cosmosdbgremlingraph


type CosmosdbGremlinGraphIndexPolicyCompositeIndexIndex struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.15.0/docs/resources/cosmosdb_gremlin_graph#order CosmosdbGremlinGraph#order}.
	Order *string `field:"required" json:"order" yaml:"order"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.15.0/docs/resources/cosmosdb_gremlin_graph#path CosmosdbGremlinGraph#path}.
	Path *string `field:"required" json:"path" yaml:"path"`
}

