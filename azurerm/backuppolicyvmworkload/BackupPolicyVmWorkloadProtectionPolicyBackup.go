// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package backuppolicyvmworkload


type BackupPolicyVmWorkloadProtectionPolicyBackup struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.15.0/docs/resources/backup_policy_vm_workload#frequency BackupPolicyVmWorkload#frequency}.
	Frequency *string `field:"optional" json:"frequency" yaml:"frequency"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.15.0/docs/resources/backup_policy_vm_workload#frequency_in_minutes BackupPolicyVmWorkload#frequency_in_minutes}.
	FrequencyInMinutes *float64 `field:"optional" json:"frequencyInMinutes" yaml:"frequencyInMinutes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.15.0/docs/resources/backup_policy_vm_workload#time BackupPolicyVmWorkload#time}.
	Time *string `field:"optional" json:"time" yaml:"time"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.15.0/docs/resources/backup_policy_vm_workload#weekdays BackupPolicyVmWorkload#weekdays}.
	Weekdays *[]*string `field:"optional" json:"weekdays" yaml:"weekdays"`
}

