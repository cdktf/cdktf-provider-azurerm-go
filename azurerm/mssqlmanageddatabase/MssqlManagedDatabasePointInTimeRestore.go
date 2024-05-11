// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mssqlmanageddatabase


type MssqlManagedDatabasePointInTimeRestore struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/mssql_managed_database#restore_point_in_time MssqlManagedDatabase#restore_point_in_time}.
	RestorePointInTime *string `field:"required" json:"restorePointInTime" yaml:"restorePointInTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/mssql_managed_database#source_database_id MssqlManagedDatabase#source_database_id}.
	SourceDatabaseId *string `field:"required" json:"sourceDatabaseId" yaml:"sourceDatabaseId"`
}

