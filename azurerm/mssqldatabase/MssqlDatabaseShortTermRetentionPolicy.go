// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mssqldatabase


type MssqlDatabaseShortTermRetentionPolicy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.111.0/docs/resources/mssql_database#retention_days MssqlDatabase#retention_days}.
	RetentionDays *float64 `field:"required" json:"retentionDays" yaml:"retentionDays"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.111.0/docs/resources/mssql_database#backup_interval_in_hours MssqlDatabase#backup_interval_in_hours}.
	BackupIntervalInHours *float64 `field:"optional" json:"backupIntervalInHours" yaml:"backupIntervalInHours"`
}

