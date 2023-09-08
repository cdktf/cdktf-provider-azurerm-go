// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mssqldatabase


type MssqlDatabaseLongTermRetentionPolicy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.72.0/docs/resources/mssql_database#monthly_retention MssqlDatabase#monthly_retention}.
	MonthlyRetention *string `field:"optional" json:"monthlyRetention" yaml:"monthlyRetention"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.72.0/docs/resources/mssql_database#weekly_retention MssqlDatabase#weekly_retention}.
	WeeklyRetention *string `field:"optional" json:"weeklyRetention" yaml:"weeklyRetention"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.72.0/docs/resources/mssql_database#week_of_year MssqlDatabase#week_of_year}.
	WeekOfYear *float64 `field:"optional" json:"weekOfYear" yaml:"weekOfYear"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.72.0/docs/resources/mssql_database#yearly_retention MssqlDatabase#yearly_retention}.
	YearlyRetention *string `field:"optional" json:"yearlyRetention" yaml:"yearlyRetention"`
}

