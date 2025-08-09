// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package postgresqlflexibleserverbackup


type PostgresqlFlexibleServerBackupTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.39.0/docs/resources/postgresql_flexible_server_backup#create PostgresqlFlexibleServerBackup#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.39.0/docs/resources/postgresql_flexible_server_backup#delete PostgresqlFlexibleServerBackup#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.39.0/docs/resources/postgresql_flexible_server_backup#read PostgresqlFlexibleServerBackup#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}

