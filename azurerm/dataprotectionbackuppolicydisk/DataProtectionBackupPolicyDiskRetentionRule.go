// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataprotectionbackuppolicydisk


type DataProtectionBackupPolicyDiskRetentionRule struct {
	// criteria block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.1.0/docs/resources/data_protection_backup_policy_disk#criteria DataProtectionBackupPolicyDisk#criteria}
	Criteria *DataProtectionBackupPolicyDiskRetentionRuleCriteria `field:"required" json:"criteria" yaml:"criteria"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.1.0/docs/resources/data_protection_backup_policy_disk#duration DataProtectionBackupPolicyDisk#duration}.
	Duration *string `field:"required" json:"duration" yaml:"duration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.1.0/docs/resources/data_protection_backup_policy_disk#name DataProtectionBackupPolicyDisk#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.1.0/docs/resources/data_protection_backup_policy_disk#priority DataProtectionBackupPolicyDisk#priority}.
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
}

