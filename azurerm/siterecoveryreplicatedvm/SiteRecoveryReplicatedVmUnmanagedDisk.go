// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package siterecoveryreplicatedvm


type SiteRecoveryReplicatedVmUnmanagedDisk struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.1.0/docs/resources/site_recovery_replicated_vm#disk_uri SiteRecoveryReplicatedVm#disk_uri}.
	DiskUri *string `field:"optional" json:"diskUri" yaml:"diskUri"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.1.0/docs/resources/site_recovery_replicated_vm#staging_storage_account_id SiteRecoveryReplicatedVm#staging_storage_account_id}.
	StagingStorageAccountId *string `field:"optional" json:"stagingStorageAccountId" yaml:"stagingStorageAccountId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.1.0/docs/resources/site_recovery_replicated_vm#target_storage_account_id SiteRecoveryReplicatedVm#target_storage_account_id}.
	TargetStorageAccountId *string `field:"optional" json:"targetStorageAccountId" yaml:"targetStorageAccountId"`
}

