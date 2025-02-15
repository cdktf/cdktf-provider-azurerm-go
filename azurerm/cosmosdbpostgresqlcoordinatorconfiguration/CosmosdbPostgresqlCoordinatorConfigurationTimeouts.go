// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cosmosdbpostgresqlcoordinatorconfiguration


type CosmosdbPostgresqlCoordinatorConfigurationTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.19.0/docs/resources/cosmosdb_postgresql_coordinator_configuration#create CosmosdbPostgresqlCoordinatorConfiguration#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.19.0/docs/resources/cosmosdb_postgresql_coordinator_configuration#delete CosmosdbPostgresqlCoordinatorConfiguration#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.19.0/docs/resources/cosmosdb_postgresql_coordinator_configuration#read CosmosdbPostgresqlCoordinatorConfiguration#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.19.0/docs/resources/cosmosdb_postgresql_coordinator_configuration#update CosmosdbPostgresqlCoordinatorConfiguration#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

