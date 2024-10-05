// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cosmosdbmongoroledefinition


type CosmosdbMongoRoleDefinitionPrivilegeResource struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.4.0/docs/resources/cosmosdb_mongo_role_definition#collection_name CosmosdbMongoRoleDefinition#collection_name}.
	CollectionName *string `field:"optional" json:"collectionName" yaml:"collectionName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.4.0/docs/resources/cosmosdb_mongo_role_definition#db_name CosmosdbMongoRoleDefinition#db_name}.
	DbName *string `field:"optional" json:"dbName" yaml:"dbName"`
}

