// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cosmosdbmongoroledefinition


type CosmosdbMongoRoleDefinitionPrivilege struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.4.0/docs/resources/cosmosdb_mongo_role_definition#actions CosmosdbMongoRoleDefinition#actions}.
	Actions *[]*string `field:"required" json:"actions" yaml:"actions"`
	// resource block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.4.0/docs/resources/cosmosdb_mongo_role_definition#resource CosmosdbMongoRoleDefinition#resource}
	Resource *CosmosdbMongoRoleDefinitionPrivilegeResource `field:"required" json:"resource" yaml:"resource"`
}

