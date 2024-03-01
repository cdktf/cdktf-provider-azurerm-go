// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package backuppolicyfileshare


type BackupPolicyFileShareRetentionMonthly struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.94.0/docs/resources/backup_policy_file_share#count BackupPolicyFileShare#count}.
	Count *float64 `field:"required" json:"count" yaml:"count"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.94.0/docs/resources/backup_policy_file_share#days BackupPolicyFileShare#days}.
	Days *[]*float64 `field:"optional" json:"days" yaml:"days"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.94.0/docs/resources/backup_policy_file_share#include_last_days BackupPolicyFileShare#include_last_days}.
	IncludeLastDays interface{} `field:"optional" json:"includeLastDays" yaml:"includeLastDays"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.94.0/docs/resources/backup_policy_file_share#weekdays BackupPolicyFileShare#weekdays}.
	Weekdays *[]*string `field:"optional" json:"weekdays" yaml:"weekdays"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.94.0/docs/resources/backup_policy_file_share#weeks BackupPolicyFileShare#weeks}.
	Weeks *[]*string `field:"optional" json:"weeks" yaml:"weeks"`
}

