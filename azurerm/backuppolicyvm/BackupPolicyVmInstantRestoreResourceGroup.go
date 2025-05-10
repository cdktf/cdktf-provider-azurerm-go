// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package backuppolicyvm


type BackupPolicyVmInstantRestoreResourceGroup struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.28.0/docs/resources/backup_policy_vm#prefix BackupPolicyVm#prefix}.
	Prefix *string `field:"required" json:"prefix" yaml:"prefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.28.0/docs/resources/backup_policy_vm#suffix BackupPolicyVm#suffix}.
	Suffix *string `field:"optional" json:"suffix" yaml:"suffix"`
}

