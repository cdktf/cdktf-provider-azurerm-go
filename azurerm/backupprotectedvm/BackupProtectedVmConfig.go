// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package backupprotectedvm

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BackupProtectedVmConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/backup_protected_vm#recovery_vault_name BackupProtectedVm#recovery_vault_name}.
	RecoveryVaultName *string `field:"required" json:"recoveryVaultName" yaml:"recoveryVaultName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/backup_protected_vm#resource_group_name BackupProtectedVm#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/backup_protected_vm#backup_policy_id BackupProtectedVm#backup_policy_id}.
	BackupPolicyId *string `field:"optional" json:"backupPolicyId" yaml:"backupPolicyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/backup_protected_vm#exclude_disk_luns BackupProtectedVm#exclude_disk_luns}.
	ExcludeDiskLuns *[]*float64 `field:"optional" json:"excludeDiskLuns" yaml:"excludeDiskLuns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/backup_protected_vm#id BackupProtectedVm#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/backup_protected_vm#include_disk_luns BackupProtectedVm#include_disk_luns}.
	IncludeDiskLuns *[]*float64 `field:"optional" json:"includeDiskLuns" yaml:"includeDiskLuns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/backup_protected_vm#protection_state BackupProtectedVm#protection_state}.
	ProtectionState *string `field:"optional" json:"protectionState" yaml:"protectionState"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/backup_protected_vm#source_vm_id BackupProtectedVm#source_vm_id}.
	SourceVmId *string `field:"optional" json:"sourceVmId" yaml:"sourceVmId"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/backup_protected_vm#timeouts BackupProtectedVm#timeouts}
	Timeouts *BackupProtectedVmTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

