// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cosmosdbmongouserdefinition


type CosmosdbMongoUserDefinitionTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.12.0/docs/resources/cosmosdb_mongo_user_definition#create CosmosdbMongoUserDefinition#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.12.0/docs/resources/cosmosdb_mongo_user_definition#delete CosmosdbMongoUserDefinition#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.12.0/docs/resources/cosmosdb_mongo_user_definition#read CosmosdbMongoUserDefinition#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.12.0/docs/resources/cosmosdb_mongo_user_definition#update CosmosdbMongoUserDefinition#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

