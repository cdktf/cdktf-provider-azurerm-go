// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataprotectionbackuppolicyblobstorage


type DataProtectionBackupPolicyBlobStorageTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.101.0/docs/resources/data_protection_backup_policy_blob_storage#create DataProtectionBackupPolicyBlobStorage#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.101.0/docs/resources/data_protection_backup_policy_blob_storage#delete DataProtectionBackupPolicyBlobStorage#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.101.0/docs/resources/data_protection_backup_policy_blob_storage#read DataProtectionBackupPolicyBlobStorage#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}

