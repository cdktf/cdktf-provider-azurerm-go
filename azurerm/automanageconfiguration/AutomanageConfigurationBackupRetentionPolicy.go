// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package automanageconfiguration


type AutomanageConfigurationBackupRetentionPolicy struct {
	// daily_schedule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/automanage_configuration#daily_schedule AutomanageConfiguration#daily_schedule}
	DailySchedule *AutomanageConfigurationBackupRetentionPolicyDailySchedule `field:"optional" json:"dailySchedule" yaml:"dailySchedule"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/automanage_configuration#retention_policy_type AutomanageConfiguration#retention_policy_type}.
	RetentionPolicyType *string `field:"optional" json:"retentionPolicyType" yaml:"retentionPolicyType"`
	// weekly_schedule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/automanage_configuration#weekly_schedule AutomanageConfiguration#weekly_schedule}
	WeeklySchedule *AutomanageConfigurationBackupRetentionPolicyWeeklySchedule `field:"optional" json:"weeklySchedule" yaml:"weeklySchedule"`
}

