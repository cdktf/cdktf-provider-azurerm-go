// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package backuppolicyvm


type BackupPolicyVmRetentionMonthly struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/backup_policy_vm#count BackupPolicyVm#count}.
	Count *float64 `field:"required" json:"count" yaml:"count"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/backup_policy_vm#days BackupPolicyVm#days}.
	Days *[]*float64 `field:"optional" json:"days" yaml:"days"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/backup_policy_vm#include_last_days BackupPolicyVm#include_last_days}.
	IncludeLastDays interface{} `field:"optional" json:"includeLastDays" yaml:"includeLastDays"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/backup_policy_vm#weekdays BackupPolicyVm#weekdays}.
	Weekdays *[]*string `field:"optional" json:"weekdays" yaml:"weekdays"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/backup_policy_vm#weeks BackupPolicyVm#weeks}.
	Weeks *[]*string `field:"optional" json:"weeks" yaml:"weeks"`
}

