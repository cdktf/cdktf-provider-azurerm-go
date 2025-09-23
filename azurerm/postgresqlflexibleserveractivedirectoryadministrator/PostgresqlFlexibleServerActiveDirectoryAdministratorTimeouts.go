// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package postgresqlflexibleserveractivedirectoryadministrator


type PostgresqlFlexibleServerActiveDirectoryAdministratorTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/postgresql_flexible_server_active_directory_administrator#create PostgresqlFlexibleServerActiveDirectoryAdministrator#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/postgresql_flexible_server_active_directory_administrator#delete PostgresqlFlexibleServerActiveDirectoryAdministrator#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/postgresql_flexible_server_active_directory_administrator#read PostgresqlFlexibleServerActiveDirectoryAdministrator#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}

