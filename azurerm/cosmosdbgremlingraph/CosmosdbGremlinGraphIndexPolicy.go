// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cosmosdbgremlingraph


type CosmosdbGremlinGraphIndexPolicy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/cosmosdb_gremlin_graph#indexing_mode CosmosdbGremlinGraph#indexing_mode}.
	IndexingMode *string `field:"required" json:"indexingMode" yaml:"indexingMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/cosmosdb_gremlin_graph#automatic CosmosdbGremlinGraph#automatic}.
	Automatic interface{} `field:"optional" json:"automatic" yaml:"automatic"`
	// composite_index block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/cosmosdb_gremlin_graph#composite_index CosmosdbGremlinGraph#composite_index}
	CompositeIndex interface{} `field:"optional" json:"compositeIndex" yaml:"compositeIndex"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/cosmosdb_gremlin_graph#excluded_paths CosmosdbGremlinGraph#excluded_paths}.
	ExcludedPaths *[]*string `field:"optional" json:"excludedPaths" yaml:"excludedPaths"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/cosmosdb_gremlin_graph#included_paths CosmosdbGremlinGraph#included_paths}.
	IncludedPaths *[]*string `field:"optional" json:"includedPaths" yaml:"includedPaths"`
	// spatial_index block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/cosmosdb_gremlin_graph#spatial_index CosmosdbGremlinGraph#spatial_index}
	SpatialIndex interface{} `field:"optional" json:"spatialIndex" yaml:"spatialIndex"`
}

