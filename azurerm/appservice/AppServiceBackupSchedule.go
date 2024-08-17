// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appservice


type AppServiceBackupSchedule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/app_service#frequency_interval AppService#frequency_interval}.
	FrequencyInterval *float64 `field:"required" json:"frequencyInterval" yaml:"frequencyInterval"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/app_service#frequency_unit AppService#frequency_unit}.
	FrequencyUnit *string `field:"required" json:"frequencyUnit" yaml:"frequencyUnit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/app_service#keep_at_least_one_backup AppService#keep_at_least_one_backup}.
	KeepAtLeastOneBackup interface{} `field:"optional" json:"keepAtLeastOneBackup" yaml:"keepAtLeastOneBackup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/app_service#retention_period_in_days AppService#retention_period_in_days}.
	RetentionPeriodInDays *float64 `field:"optional" json:"retentionPeriodInDays" yaml:"retentionPeriodInDays"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/app_service#start_time AppService#start_time}.
	StartTime *string `field:"optional" json:"startTime" yaml:"startTime"`
}

