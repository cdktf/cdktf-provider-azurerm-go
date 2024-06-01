// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mssqlmanageddatabase


type MssqlManagedDatabaseLongTermRetentionPolicy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.106.1/docs/resources/mssql_managed_database#immutable_backups_enabled MssqlManagedDatabase#immutable_backups_enabled}.
	ImmutableBackupsEnabled interface{} `field:"optional" json:"immutableBackupsEnabled" yaml:"immutableBackupsEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.106.1/docs/resources/mssql_managed_database#monthly_retention MssqlManagedDatabase#monthly_retention}.
	MonthlyRetention *string `field:"optional" json:"monthlyRetention" yaml:"monthlyRetention"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.106.1/docs/resources/mssql_managed_database#weekly_retention MssqlManagedDatabase#weekly_retention}.
	WeeklyRetention *string `field:"optional" json:"weeklyRetention" yaml:"weeklyRetention"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.106.1/docs/resources/mssql_managed_database#week_of_year MssqlManagedDatabase#week_of_year}.
	WeekOfYear *float64 `field:"optional" json:"weekOfYear" yaml:"weekOfYear"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.106.1/docs/resources/mssql_managed_database#yearly_retention MssqlManagedDatabase#yearly_retention}.
	YearlyRetention *string `field:"optional" json:"yearlyRetention" yaml:"yearlyRetention"`
}

