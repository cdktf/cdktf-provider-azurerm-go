// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cosmosdbsqlcontainer


type CosmosdbSqlContainerIndexingPolicy struct {
	// composite_index block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/cosmosdb_sql_container#composite_index CosmosdbSqlContainer#composite_index}
	CompositeIndex interface{} `field:"optional" json:"compositeIndex" yaml:"compositeIndex"`
	// excluded_path block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/cosmosdb_sql_container#excluded_path CosmosdbSqlContainer#excluded_path}
	ExcludedPath interface{} `field:"optional" json:"excludedPath" yaml:"excludedPath"`
	// included_path block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/cosmosdb_sql_container#included_path CosmosdbSqlContainer#included_path}
	IncludedPath interface{} `field:"optional" json:"includedPath" yaml:"includedPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/cosmosdb_sql_container#indexing_mode CosmosdbSqlContainer#indexing_mode}.
	IndexingMode *string `field:"optional" json:"indexingMode" yaml:"indexingMode"`
	// spatial_index block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/cosmosdb_sql_container#spatial_index CosmosdbSqlContainer#spatial_index}
	SpatialIndex interface{} `field:"optional" json:"spatialIndex" yaml:"spatialIndex"`
}

