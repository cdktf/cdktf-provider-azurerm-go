// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cosmosdbpostgresqlrole


type CosmosdbPostgresqlRoleTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/cosmosdb_postgresql_role#create CosmosdbPostgresqlRole#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/cosmosdb_postgresql_role#delete CosmosdbPostgresqlRole#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/cosmosdb_postgresql_role#read CosmosdbPostgresqlRole#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}

