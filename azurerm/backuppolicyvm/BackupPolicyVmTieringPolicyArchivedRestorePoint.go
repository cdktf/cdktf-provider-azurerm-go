// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package backuppolicyvm


type BackupPolicyVmTieringPolicyArchivedRestorePoint struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.17.0/docs/resources/backup_policy_vm#mode BackupPolicyVm#mode}.
	Mode *string `field:"required" json:"mode" yaml:"mode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.17.0/docs/resources/backup_policy_vm#duration BackupPolicyVm#duration}.
	Duration *float64 `field:"optional" json:"duration" yaml:"duration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.17.0/docs/resources/backup_policy_vm#duration_type BackupPolicyVm#duration_type}.
	DurationType *string `field:"optional" json:"durationType" yaml:"durationType"`
}

