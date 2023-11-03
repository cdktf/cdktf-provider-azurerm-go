// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package windowswebapp


type WindowsWebAppBackupSchedule struct {
	// How often the backup should be executed (e.g. for weekly backup, this should be set to `7` and `frequency_unit` should be set to `Day`).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.79.0/docs/resources/windows_web_app#frequency_interval WindowsWebApp#frequency_interval}
	FrequencyInterval *float64 `field:"required" json:"frequencyInterval" yaml:"frequencyInterval"`
	// The unit of time for how often the backup should take place. Possible values include: `Day` and `Hour`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.79.0/docs/resources/windows_web_app#frequency_unit WindowsWebApp#frequency_unit}
	FrequencyUnit *string `field:"required" json:"frequencyUnit" yaml:"frequencyUnit"`
	// Should the service keep at least one backup, regardless of age of backup. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.79.0/docs/resources/windows_web_app#keep_at_least_one_backup WindowsWebApp#keep_at_least_one_backup}
	KeepAtLeastOneBackup interface{} `field:"optional" json:"keepAtLeastOneBackup" yaml:"keepAtLeastOneBackup"`
	// After how many days backups should be deleted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.79.0/docs/resources/windows_web_app#retention_period_days WindowsWebApp#retention_period_days}
	RetentionPeriodDays *float64 `field:"optional" json:"retentionPeriodDays" yaml:"retentionPeriodDays"`
	// When the schedule should start working in RFC-3339 format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.79.0/docs/resources/windows_web_app#start_time WindowsWebApp#start_time}
	StartTime *string `field:"optional" json:"startTime" yaml:"startTime"`
}

