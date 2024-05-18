// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataprotectionbackuppolicydisk


type DataProtectionBackupPolicyDiskTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.0/docs/resources/data_protection_backup_policy_disk#create DataProtectionBackupPolicyDisk#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.0/docs/resources/data_protection_backup_policy_disk#delete DataProtectionBackupPolicyDisk#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.0/docs/resources/data_protection_backup_policy_disk#read DataProtectionBackupPolicyDisk#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}

