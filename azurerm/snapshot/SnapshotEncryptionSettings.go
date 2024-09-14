// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package snapshot


type SnapshotEncryptionSettings struct {
	// disk_encryption_key block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.2.0/docs/resources/snapshot#disk_encryption_key Snapshot#disk_encryption_key}
	DiskEncryptionKey *SnapshotEncryptionSettingsDiskEncryptionKey `field:"required" json:"diskEncryptionKey" yaml:"diskEncryptionKey"`
	// key_encryption_key block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.2.0/docs/resources/snapshot#key_encryption_key Snapshot#key_encryption_key}
	KeyEncryptionKey *SnapshotEncryptionSettingsKeyEncryptionKey `field:"optional" json:"keyEncryptionKey" yaml:"keyEncryptionKey"`
}

