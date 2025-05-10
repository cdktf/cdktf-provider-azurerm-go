// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package windowsvirtualmachinescaleset


type WindowsVirtualMachineScaleSetDataDisk struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.28.0/docs/resources/windows_virtual_machine_scale_set#caching WindowsVirtualMachineScaleSet#caching}.
	Caching *string `field:"required" json:"caching" yaml:"caching"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.28.0/docs/resources/windows_virtual_machine_scale_set#disk_size_gb WindowsVirtualMachineScaleSet#disk_size_gb}.
	DiskSizeGb *float64 `field:"required" json:"diskSizeGb" yaml:"diskSizeGb"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.28.0/docs/resources/windows_virtual_machine_scale_set#lun WindowsVirtualMachineScaleSet#lun}.
	Lun *float64 `field:"required" json:"lun" yaml:"lun"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.28.0/docs/resources/windows_virtual_machine_scale_set#storage_account_type WindowsVirtualMachineScaleSet#storage_account_type}.
	StorageAccountType *string `field:"required" json:"storageAccountType" yaml:"storageAccountType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.28.0/docs/resources/windows_virtual_machine_scale_set#create_option WindowsVirtualMachineScaleSet#create_option}.
	CreateOption *string `field:"optional" json:"createOption" yaml:"createOption"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.28.0/docs/resources/windows_virtual_machine_scale_set#disk_encryption_set_id WindowsVirtualMachineScaleSet#disk_encryption_set_id}.
	DiskEncryptionSetId *string `field:"optional" json:"diskEncryptionSetId" yaml:"diskEncryptionSetId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.28.0/docs/resources/windows_virtual_machine_scale_set#name WindowsVirtualMachineScaleSet#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.28.0/docs/resources/windows_virtual_machine_scale_set#ultra_ssd_disk_iops_read_write WindowsVirtualMachineScaleSet#ultra_ssd_disk_iops_read_write}.
	UltraSsdDiskIopsReadWrite *float64 `field:"optional" json:"ultraSsdDiskIopsReadWrite" yaml:"ultraSsdDiskIopsReadWrite"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.28.0/docs/resources/windows_virtual_machine_scale_set#ultra_ssd_disk_mbps_read_write WindowsVirtualMachineScaleSet#ultra_ssd_disk_mbps_read_write}.
	UltraSsdDiskMbpsReadWrite *float64 `field:"optional" json:"ultraSsdDiskMbpsReadWrite" yaml:"ultraSsdDiskMbpsReadWrite"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.28.0/docs/resources/windows_virtual_machine_scale_set#write_accelerator_enabled WindowsVirtualMachineScaleSet#write_accelerator_enabled}.
	WriteAcceleratorEnabled interface{} `field:"optional" json:"writeAcceleratorEnabled" yaml:"writeAcceleratorEnabled"`
}

