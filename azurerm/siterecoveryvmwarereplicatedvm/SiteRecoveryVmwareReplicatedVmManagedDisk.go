// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package siterecoveryvmwarereplicatedvm


type SiteRecoveryVmwareReplicatedVmManagedDisk struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/site_recovery_vmware_replicated_vm#disk_id SiteRecoveryVmwareReplicatedVm#disk_id}.
	DiskId *string `field:"required" json:"diskId" yaml:"diskId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/site_recovery_vmware_replicated_vm#target_disk_type SiteRecoveryVmwareReplicatedVm#target_disk_type}.
	TargetDiskType *string `field:"required" json:"targetDiskType" yaml:"targetDiskType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/site_recovery_vmware_replicated_vm#log_storage_account_id SiteRecoveryVmwareReplicatedVm#log_storage_account_id}.
	LogStorageAccountId *string `field:"optional" json:"logStorageAccountId" yaml:"logStorageAccountId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/site_recovery_vmware_replicated_vm#target_disk_encryption_set_id SiteRecoveryVmwareReplicatedVm#target_disk_encryption_set_id}.
	TargetDiskEncryptionSetId *string `field:"optional" json:"targetDiskEncryptionSetId" yaml:"targetDiskEncryptionSetId"`
}

