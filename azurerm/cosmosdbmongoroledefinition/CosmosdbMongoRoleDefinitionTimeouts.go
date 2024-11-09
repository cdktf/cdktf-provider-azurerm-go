// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cosmosdbmongoroledefinition


type CosmosdbMongoRoleDefinitionTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.9.0/docs/resources/cosmosdb_mongo_role_definition#create CosmosdbMongoRoleDefinition#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.9.0/docs/resources/cosmosdb_mongo_role_definition#delete CosmosdbMongoRoleDefinition#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.9.0/docs/resources/cosmosdb_mongo_role_definition#read CosmosdbMongoRoleDefinition#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.9.0/docs/resources/cosmosdb_mongo_role_definition#update CosmosdbMongoRoleDefinition#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

