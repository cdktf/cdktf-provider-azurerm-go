// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package postgresqlactivedirectoryadministrator


type PostgresqlActiveDirectoryAdministratorTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/postgresql_active_directory_administrator#create PostgresqlActiveDirectoryAdministrator#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/postgresql_active_directory_administrator#delete PostgresqlActiveDirectoryAdministrator#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/postgresql_active_directory_administrator#read PostgresqlActiveDirectoryAdministrator#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/postgresql_active_directory_administrator#update PostgresqlActiveDirectoryAdministrator#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

