// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package windowsvirtualmachine


type WindowsVirtualMachineOsDisk struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs/resources/windows_virtual_machine#caching WindowsVirtualMachine#caching}.
	Caching *string `field:"required" json:"caching" yaml:"caching"`
	// diff_disk_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs/resources/windows_virtual_machine#diff_disk_settings WindowsVirtualMachine#diff_disk_settings}
	DiffDiskSettings *WindowsVirtualMachineOsDiskDiffDiskSettings `field:"optional" json:"diffDiskSettings" yaml:"diffDiskSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs/resources/windows_virtual_machine#disk_encryption_set_id WindowsVirtualMachine#disk_encryption_set_id}.
	DiskEncryptionSetId *string `field:"optional" json:"diskEncryptionSetId" yaml:"diskEncryptionSetId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs/resources/windows_virtual_machine#disk_size_gb WindowsVirtualMachine#disk_size_gb}.
	DiskSizeGb *float64 `field:"optional" json:"diskSizeGb" yaml:"diskSizeGb"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs/resources/windows_virtual_machine#name WindowsVirtualMachine#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs/resources/windows_virtual_machine#secure_vm_disk_encryption_set_id WindowsVirtualMachine#secure_vm_disk_encryption_set_id}.
	SecureVmDiskEncryptionSetId *string `field:"optional" json:"secureVmDiskEncryptionSetId" yaml:"secureVmDiskEncryptionSetId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs/resources/windows_virtual_machine#security_encryption_type WindowsVirtualMachine#security_encryption_type}.
	SecurityEncryptionType *string `field:"optional" json:"securityEncryptionType" yaml:"securityEncryptionType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs/resources/windows_virtual_machine#storage_account_type WindowsVirtualMachine#storage_account_type}.
	StorageAccountType *string `field:"optional" json:"storageAccountType" yaml:"storageAccountType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs/resources/windows_virtual_machine#write_accelerator_enabled WindowsVirtualMachine#write_accelerator_enabled}.
	WriteAcceleratorEnabled interface{} `field:"optional" json:"writeAcceleratorEnabled" yaml:"writeAcceleratorEnabled"`
}

