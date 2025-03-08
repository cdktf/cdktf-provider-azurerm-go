// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataprotectionbackuppolicyblobstorage


type DataProtectionBackupPolicyBlobStorageRetentionRuleCriteria struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.22.0/docs/resources/data_protection_backup_policy_blob_storage#absolute_criteria DataProtectionBackupPolicyBlobStorage#absolute_criteria}.
	AbsoluteCriteria *string `field:"optional" json:"absoluteCriteria" yaml:"absoluteCriteria"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.22.0/docs/resources/data_protection_backup_policy_blob_storage#days_of_month DataProtectionBackupPolicyBlobStorage#days_of_month}.
	DaysOfMonth *[]*float64 `field:"optional" json:"daysOfMonth" yaml:"daysOfMonth"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.22.0/docs/resources/data_protection_backup_policy_blob_storage#days_of_week DataProtectionBackupPolicyBlobStorage#days_of_week}.
	DaysOfWeek *[]*string `field:"optional" json:"daysOfWeek" yaml:"daysOfWeek"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.22.0/docs/resources/data_protection_backup_policy_blob_storage#months_of_year DataProtectionBackupPolicyBlobStorage#months_of_year}.
	MonthsOfYear *[]*string `field:"optional" json:"monthsOfYear" yaml:"monthsOfYear"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.22.0/docs/resources/data_protection_backup_policy_blob_storage#scheduled_backup_times DataProtectionBackupPolicyBlobStorage#scheduled_backup_times}.
	ScheduledBackupTimes *[]*string `field:"optional" json:"scheduledBackupTimes" yaml:"scheduledBackupTimes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.22.0/docs/resources/data_protection_backup_policy_blob_storage#weeks_of_month DataProtectionBackupPolicyBlobStorage#weeks_of_month}.
	WeeksOfMonth *[]*string `field:"optional" json:"weeksOfMonth" yaml:"weeksOfMonth"`
}

