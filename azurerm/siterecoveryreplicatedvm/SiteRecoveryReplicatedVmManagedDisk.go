// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package siterecoveryreplicatedvm


type SiteRecoveryReplicatedVmManagedDisk struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.17.0/docs/resources/site_recovery_replicated_vm#disk_id SiteRecoveryReplicatedVm#disk_id}.
	DiskId *string `field:"optional" json:"diskId" yaml:"diskId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.17.0/docs/resources/site_recovery_replicated_vm#staging_storage_account_id SiteRecoveryReplicatedVm#staging_storage_account_id}.
	StagingStorageAccountId *string `field:"optional" json:"stagingStorageAccountId" yaml:"stagingStorageAccountId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.17.0/docs/resources/site_recovery_replicated_vm#target_disk_encryption SiteRecoveryReplicatedVm#target_disk_encryption}.
	TargetDiskEncryption interface{} `field:"optional" json:"targetDiskEncryption" yaml:"targetDiskEncryption"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.17.0/docs/resources/site_recovery_replicated_vm#target_disk_encryption_set_id SiteRecoveryReplicatedVm#target_disk_encryption_set_id}.
	TargetDiskEncryptionSetId *string `field:"optional" json:"targetDiskEncryptionSetId" yaml:"targetDiskEncryptionSetId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.17.0/docs/resources/site_recovery_replicated_vm#target_disk_type SiteRecoveryReplicatedVm#target_disk_type}.
	TargetDiskType *string `field:"optional" json:"targetDiskType" yaml:"targetDiskType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.17.0/docs/resources/site_recovery_replicated_vm#target_replica_disk_type SiteRecoveryReplicatedVm#target_replica_disk_type}.
	TargetReplicaDiskType *string `field:"optional" json:"targetReplicaDiskType" yaml:"targetReplicaDiskType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.17.0/docs/resources/site_recovery_replicated_vm#target_resource_group_id SiteRecoveryReplicatedVm#target_resource_group_id}.
	TargetResourceGroupId *string `field:"optional" json:"targetResourceGroupId" yaml:"targetResourceGroupId"`
}

