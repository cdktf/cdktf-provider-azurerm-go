// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataprotectionbackuppolicymysqlflexibleserver


type DataProtectionBackupPolicyMysqlFlexibleServerRetentionRuleCriteria struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.33.0/docs/resources/data_protection_backup_policy_mysql_flexible_server#absolute_criteria DataProtectionBackupPolicyMysqlFlexibleServer#absolute_criteria}.
	AbsoluteCriteria *string `field:"optional" json:"absoluteCriteria" yaml:"absoluteCriteria"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.33.0/docs/resources/data_protection_backup_policy_mysql_flexible_server#days_of_week DataProtectionBackupPolicyMysqlFlexibleServer#days_of_week}.
	DaysOfWeek *[]*string `field:"optional" json:"daysOfWeek" yaml:"daysOfWeek"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.33.0/docs/resources/data_protection_backup_policy_mysql_flexible_server#months_of_year DataProtectionBackupPolicyMysqlFlexibleServer#months_of_year}.
	MonthsOfYear *[]*string `field:"optional" json:"monthsOfYear" yaml:"monthsOfYear"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.33.0/docs/resources/data_protection_backup_policy_mysql_flexible_server#scheduled_backup_times DataProtectionBackupPolicyMysqlFlexibleServer#scheduled_backup_times}.
	ScheduledBackupTimes *[]*string `field:"optional" json:"scheduledBackupTimes" yaml:"scheduledBackupTimes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.33.0/docs/resources/data_protection_backup_policy_mysql_flexible_server#weeks_of_month DataProtectionBackupPolicyMysqlFlexibleServer#weeks_of_month}.
	WeeksOfMonth *[]*string `field:"optional" json:"weeksOfMonth" yaml:"weeksOfMonth"`
}

