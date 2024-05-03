// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataprotectionbackuppolicykubernetescluster


type DataProtectionBackupPolicyKubernetesClusterRetentionRuleCriteria struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.102.0/docs/resources/data_protection_backup_policy_kubernetes_cluster#absolute_criteria DataProtectionBackupPolicyKubernetesCluster#absolute_criteria}.
	AbsoluteCriteria *string `field:"optional" json:"absoluteCriteria" yaml:"absoluteCriteria"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.102.0/docs/resources/data_protection_backup_policy_kubernetes_cluster#days_of_week DataProtectionBackupPolicyKubernetesCluster#days_of_week}.
	DaysOfWeek *[]*string `field:"optional" json:"daysOfWeek" yaml:"daysOfWeek"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.102.0/docs/resources/data_protection_backup_policy_kubernetes_cluster#months_of_year DataProtectionBackupPolicyKubernetesCluster#months_of_year}.
	MonthsOfYear *[]*string `field:"optional" json:"monthsOfYear" yaml:"monthsOfYear"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.102.0/docs/resources/data_protection_backup_policy_kubernetes_cluster#scheduled_backup_times DataProtectionBackupPolicyKubernetesCluster#scheduled_backup_times}.
	ScheduledBackupTimes *[]*string `field:"optional" json:"scheduledBackupTimes" yaml:"scheduledBackupTimes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.102.0/docs/resources/data_protection_backup_policy_kubernetes_cluster#weeks_of_month DataProtectionBackupPolicyKubernetesCluster#weeks_of_month}.
	WeeksOfMonth *[]*string `field:"optional" json:"weeksOfMonth" yaml:"weeksOfMonth"`
}

