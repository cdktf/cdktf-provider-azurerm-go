// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cosmosdbpostgresqlcluster


type CosmosdbPostgresqlClusterTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/cosmosdb_postgresql_cluster#create CosmosdbPostgresqlCluster#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/cosmosdb_postgresql_cluster#delete CosmosdbPostgresqlCluster#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/cosmosdb_postgresql_cluster#read CosmosdbPostgresqlCluster#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/cosmosdb_postgresql_cluster#update CosmosdbPostgresqlCluster#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

