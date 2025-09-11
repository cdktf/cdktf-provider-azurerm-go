// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package backuppolicyfileshare


type BackupPolicyFileShareBackupHourly struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs/resources/backup_policy_file_share#interval BackupPolicyFileShare#interval}.
	Interval *float64 `field:"required" json:"interval" yaml:"interval"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs/resources/backup_policy_file_share#start_time BackupPolicyFileShare#start_time}.
	StartTime *string `field:"required" json:"startTime" yaml:"startTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs/resources/backup_policy_file_share#window_duration BackupPolicyFileShare#window_duration}.
	WindowDuration *float64 `field:"required" json:"windowDuration" yaml:"windowDuration"`
}

