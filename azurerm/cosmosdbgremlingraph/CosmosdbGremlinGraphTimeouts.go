// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cosmosdbgremlingraph


type CosmosdbGremlinGraphTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/cosmosdb_gremlin_graph#create CosmosdbGremlinGraph#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/cosmosdb_gremlin_graph#delete CosmosdbGremlinGraph#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/cosmosdb_gremlin_graph#read CosmosdbGremlinGraph#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/cosmosdb_gremlin_graph#update CosmosdbGremlinGraph#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

