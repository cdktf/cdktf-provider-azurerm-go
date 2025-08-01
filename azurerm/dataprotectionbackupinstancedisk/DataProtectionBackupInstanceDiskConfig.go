// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataprotectionbackupinstancedisk

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataProtectionBackupInstanceDiskConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/data_protection_backup_instance_disk#backup_policy_id DataProtectionBackupInstanceDisk#backup_policy_id}.
	BackupPolicyId *string `field:"required" json:"backupPolicyId" yaml:"backupPolicyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/data_protection_backup_instance_disk#disk_id DataProtectionBackupInstanceDisk#disk_id}.
	DiskId *string `field:"required" json:"diskId" yaml:"diskId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/data_protection_backup_instance_disk#location DataProtectionBackupInstanceDisk#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/data_protection_backup_instance_disk#name DataProtectionBackupInstanceDisk#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/data_protection_backup_instance_disk#snapshot_resource_group_name DataProtectionBackupInstanceDisk#snapshot_resource_group_name}.
	SnapshotResourceGroupName *string `field:"required" json:"snapshotResourceGroupName" yaml:"snapshotResourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/data_protection_backup_instance_disk#vault_id DataProtectionBackupInstanceDisk#vault_id}.
	VaultId *string `field:"required" json:"vaultId" yaml:"vaultId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/data_protection_backup_instance_disk#id DataProtectionBackupInstanceDisk#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/data_protection_backup_instance_disk#snapshot_subscription_id DataProtectionBackupInstanceDisk#snapshot_subscription_id}.
	SnapshotSubscriptionId *string `field:"optional" json:"snapshotSubscriptionId" yaml:"snapshotSubscriptionId"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/data_protection_backup_instance_disk#timeouts DataProtectionBackupInstanceDisk#timeouts}
	Timeouts *DataProtectionBackupInstanceDiskTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

