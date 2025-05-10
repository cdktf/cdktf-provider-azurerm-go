// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cosmosdbsqlroledefinition


type CosmosdbSqlRoleDefinitionPermissions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.28.0/docs/resources/cosmosdb_sql_role_definition#data_actions CosmosdbSqlRoleDefinition#data_actions}.
	DataActions *[]*string `field:"required" json:"dataActions" yaml:"dataActions"`
}

